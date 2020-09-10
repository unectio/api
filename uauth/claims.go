/////////////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2019-2020, Unectio Inc, All Right Reserved.
//
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
/////////////////////////////////////////////////////////////////////////////////

package uauth

import (
	"time"
	"errors"
	"context"
	"github.com/unectio/db"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
)

type rawClaims struct {
	*jwt.StandardClaims
	ProjectId	string		`json:"projectid"`
	RoleName	string		`json:"role"`
	UserId		string		`json:"userid"`
}

type Claims struct {
	User		bson.ObjectId
	Project		bson.ObjectId	/* Can be empty, meaning the "default" project */
	Role		*Role		/* Can be nil, meaning "roleNobody" FIXME */
	Scope		*db.KeyScopeDb
}

func (c *rawClaims)project() (bson.ObjectId, bool) {
	if c.ProjectId != "" && c.ProjectId != db.DefaultProjectId {
		return mongo.ObjectId(c.ProjectId)
	} else {
		return "", true
	}
}

func (claims *Claims)setupServerSigned(ctx context.Context, rc *rawClaims) error {
	/*
	 * Server (us) signed claims are fully trusted, but we still need the
	 * type conversion.
	 */
	u, ok := mongo.ObjectId(rc.UserId)
	if !ok {
		return errors.New("Bad userid encoded by server\n")
	}

	p, ok := rc.project()
	if !ok {
		return errors.New("Bad projectid encoded by server\n")
	}

	claims.User = u
	claims.Project = p
	claims.Role = GetRole(rc.RoleName)

	return nil
}

func (claims *Claims)setupSelfSigned(ctx context.Context, rc *rawClaims, key *db.KeyDb) error {
	/*
	 * Self-signed claims. In this case it can only claim the project
	 * to work on and it must not be valid longer than several minutes.
	 */
	if rc.UserId != "" || rc.RoleName != "" {
		return errors.New("Cannot claim user or role")
	}

	if rc.StandardClaims.ExpiresAt == 0 ||
			rc.StandardClaims.ExpiresAt > time.Now().Add(SelfKeyLifetime + MaxTimerDesync).Unix() {
		return errors.New("Too long-living key")
	}

	var u db.UserDb

	err := db.Load(ctx, key.User, &u)
	if err != nil {
		return errors.New("Cannot get user for key\n")
	}

	var ok bool

	claims.User = u.ID()
	claims.Project, ok = rc.project()
	if !ok {
		return errors.New("Bad project claimed")
	}

	/*
	 * Not claiming a project means working with the default one
	 * and it's just allowed with the default user role.
	 */
	if claims.Project == "" {
		claims.Role = GetRole(u.Role)
		return nil
	}

	/*
	 * Admin user can claim any project without any checks.
	 * XXX maybe checking the project for existance is not a bad idea?
	 * XXX is it also worth checking the key scope?
	 * FIXME No db.XXXProject() cast here, as the project is claimed by admin
	 * by its exact name.
	 */

	role := GetRole(u.Role)
	if role.Can(RoleCapAnyProject) {
		claims.Role = role
		return nil
	}

	/*
	 * Otherwise, the user must have a role in the target project,
	 * so go ahead and find out one.
	 */

	var r db.RoleDb

	err = db.Find(ctx, bson.M{"user": u.ID(), "project": claims.Project}, &r)
	if err != nil {
		return err
	}

	claims.Role = GetRole(r.Role)
	return nil
}

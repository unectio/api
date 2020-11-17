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

package api

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	AuthTokHeader = "X-Auth-Token"
	SubjTokHeader = "X-Subject-Token"
)

const (
	ReferencePrefix    string = "$ref:"
	RefPlatform        string = "platform"
	RefAuth            string = "auth"
	RefAuthJWTSignKey  string = "jwt_sign_key"
	RefSecret          string = "secret"
	RefWebsockToken    string = "access_token"
	RefWebsockEndpoint string = "endpoint"
	RefValue           string = "valueof"
	RefMware           string = "mware"
	RefAddr            string = "address"
	RefUser            string = "user"
	RefPass            string = "password"

	PlatformDbCreds string = "db_creds"
	ValueProject    string = "project"

	AutoValue = "auto"

	URLProjectPfx = "$project."

	CodeUpdateRamped   = "ramped"
	CodeUpdateRecreate = "recreate"

	RepoDescFile = ".faas.yml"
)

const (
	URLCookieLength = 32
)

type ObjectId string

func FromDb(oid bson.ObjectId) ObjectId {
	return ObjectId(oid.Hex())
}

func (oid ObjectId) ToDb() (bson.ObjectId, bool) {
	return mongo.ObjectId(string(oid))
}

/*
 * All manipulated objects have this stuff on-board
 */
type CommonImage struct {
	Id       ObjectId `json:"id"                     yaml:"-"`
	Name     string   `json:"name"                   yaml:"name"`
	Tags     []string `json:"tags,omitempty"         yaml:"tags,omitempty"`
	UserData string   `json:"userdata,omitempty"     yaml:"-"`
}

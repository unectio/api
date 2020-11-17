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
	"context"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/unectio/db"
	"github.com/unectio/util"
	sc "github.com/unectio/util/context"
	"github.com/unectio/util/mongo"
)

func lookupSignKey(ctx context.Context, tok *jwt.Token) (*db.KeyDb, error) {
	if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected sign method: %v", tok.Header["alg"])
	}

	kid, ok := tok.Header["kid"].(string)
	if !ok {
		return nil, fmt.Errorf("Unexpected key id: %v", tok.Header["kid"])
	}

	var key db.KeyDb

	q := mongo.IdSafeQ(nil, kid)
	if q == nil {
		return nil, errors.New("Bad keyid")
	}

	err := db.Find(ctx, q, &key)
	if err != nil {
		return nil, err
	}

	return &key, nil
}

func DecodeToken(ctx context.Context, token string) *Claims {
	token, ok := util.ParseBearer(token)
	if !ok {
		return nil
	}

	var key *db.KeyDb

	tok, err := jwt.ParseWithClaims(token, &rawClaims{},
		func(tok *jwt.Token) (interface{}, error) {
			var err error

			key, err = lookupSignKey(ctx, tok)
			if err != nil {
				return nil, err
			}

			return key.Value, nil
		})

	if err != nil {
		sc.L(ctx).Errorf("Error parsing token: %s\n", err.Error())
		return nil
	}

	rc := tok.Claims.(*rawClaims)
	claims := &Claims{}

	switch key.Kind {
	case db.KeyKindSelf:
		err = claims.setupSelfSigned(ctx, rc, key)
	case db.KeyKindServer:
		err = claims.setupServerSigned(ctx, rc)
	default:
		err = fmt.Errorf("Unknown token kind (%s)", key.Kind)
	}

	if err != nil {
		sc.L(ctx).Errorf("Bad token: %s\n", err.Error())
		return nil
	}

	claims.Scope = key.Scope
	if key.Scope != nil {
		/*
		 * Key can be project-bound
		 */
		if key.Scope.Project != "" && key.Scope.Project != claims.Project {
			return nil
		}
	}

	return claims
}

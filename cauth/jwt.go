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

package cauth

import (
	"fmt"
	"context"
	"net/http"
	"github.com/unectio/db"
	"github.com/unectio/api"
	"github.com/unectio/util"
	"github.com/dgrijalva/jwt-go"
)

type jwtAuth struct {
	db.AuthJWTDb
}

func (ja *jwtAuth)Verify(ctx context.Context, w http.ResponseWriter, r *http.Request) map[string]interface{} {
	h := r.Header.Get(api.AuthCommonHeader)
	if h == "" {
		http.Error(w, "Need " + api.AuthCommonHeader + " header", http.StatusUnauthorized)
		return nil
	}

	return ja.verifyKey(ctx, h, w, r)
}

func (ja *jwtAuth)verifyKey(ctx context.Context, h string, w http.ResponseWriter, r *http.Request) map[string]interface{} {

	token, ok := util.ParseBearer(h)
	if !ok {
		http.Error(w, "Need Bearer scheme", http.StatusUnauthorized)
		return nil
	}

	tok, err := jwt.ParseWithClaims(token, jwt.MapClaims{},
			func(tok *jwt.Token) (interface{}, error) {
				if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected sign method: %v", tok.Header["alg"])
				}
				return ja.Key, nil
			})

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return nil
	}

	return tok.Claims.(jwt.MapClaims)
}

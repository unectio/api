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

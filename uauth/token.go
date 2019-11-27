package uauth

import (
	"fmt"
	"errors"
	"context"
	"github.com/unectio/db"
	"github.com/unectio/util"
	"github.com/unectio/util/mongo"
	sc "github.com/unectio/util/context"
	"github.com/dgrijalva/jwt-go"
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

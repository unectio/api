package cauth

import (
	"context"
	"net/http"
	"github.com/unectio/api"
	"github.com/unectio/api/uauth"
)

type platformAuth struct { }

func (pa *platformAuth)Verify(ctx context.Context, w http.ResponseWriter, r *http.Request) map[string]interface{} {
	h := r.Header.Get(api.AuthCommonHeader)
	if h == "" {
		http.Error(w, "Need " + api.AuthCommonHeader + " header", http.StatusUnauthorized)
		return nil
	}

	claims := uauth.DecodeToken(ctx, h)
	if claims == nil {
		http.Error(w, "", http.StatusForbidden)
		return nil
	}

	return map[string]interface{} {
		"user":		claims.User.Hex(),
		"project":	claims.Project.Hex(),
		"role":		claims.Role.Name,
		"token":	h,
	}
}


package cauth

import (
	"context"
	"net/http"
	"github.com/unectio/api"
)

type plainAuth struct {}

func (_ *plainAuth)Verify(ctx context.Context, w http.ResponseWriter, r *http.Request) map[string]interface{} {
	h := r.Header.Get(api.AuthCommonHeader)
	if h == "" {
		http.Error(w, "Need " + api.AuthCommonHeader + " header", http.StatusUnauthorized)
		return nil
	}

	return map[string]interface{} {
		api.AuthCommonHeader: h,
	}
}

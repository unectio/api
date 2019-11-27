package cauth

import (
	"errors"
	"context"
	"this/db"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

type AuthMethod interface {
	Verify(context.Context, http.ResponseWriter, *http.Request) map[string]interface{}
}

func AuthSetup(ctx context.Context, id bson.ObjectId) (AuthMethod, error) {
	var am db.AuthMethodDb

	err := db.Load(ctx, id, &am)
	if err != nil {
		return nil, errors.New("no auth method: " + err.Error())
	}

	var ami AuthMethod

	switch {
	case am.JWT != nil:
		ami = &jwtAuth{ *am.JWT}
	case am.Platform:
		ami = &platformAuth{}
	default:
		ami = &plainAuth{}
	}

	return ami, nil
}

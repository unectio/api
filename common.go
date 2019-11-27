package api

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	AuthTokHeader		= "X-Auth-Token"
	SubjTokHeader		= "X-Subject-Token"
)

const (
	ReferencePrefix	string		= "$ref:"
	RefPlatform string		= "platform"
	RefAuth string			= "auth"
	RefAuthJWTSignKey string	= "jwt_sign_key"
	RefSecret string		= "secret"
	RefWebsockToken string		= "access_token"
	RefWebsockEndpoint string	= "endpoint"
	RefValue string			= "valueof"
	RefMware string			= "mware"
	RefAddr string			= "address"
	RefUser string			= "user"
	RefPass string			= "password"

	PlatformDbCreds string		= "db_creds"
	ValueProject string		= "project"

	AutoValue			= "auto"

	URLProjectPfx			= "$project."

	CodeUpdateRamped		= "ramped"
	CodeUpdateRecreate		= "recreate"

	RepoDescFile			= ".faas.yml"
)

const (
	URLCookieLength		= 32
)

type ObjectId string

func FromDb(oid bson.ObjectId) ObjectId {
	return ObjectId(oid.Hex())
}

func (oid ObjectId)ToDb() (bson.ObjectId, bool) {
	return mongo.ObjectId(string(oid))
}

/*
 * All manipulated objects have this stuff on-board
 */
type CommonImage struct {
	Id		ObjectId		`json:"id"                     yaml:"-"`
	Name		string			`json:"name"                   yaml:"name"`
	Tags		[]string		`json:"tags,omitempty"         yaml:"tags,omitempty"`
	UserData	string			`json:"userdata,omitempty"     yaml:"-"`
}

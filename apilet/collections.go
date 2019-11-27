package apilet

import (
	"github.com/unectio/util/restmux/client"
)

var Functions		= &client.Collection{"functions", nil}
var AuthMethods		= &client.Collection{"auths", nil}
var Repos		= &client.Collection{"repositories", nil}
var Routers		= &client.Collection{"routers", nil}
var Secrets		= &client.Collection{"secrets", nil}

var FnCodes		= &client.Collection{"code", Functions}
var FnTriggers		= &client.Collection{"triggers", Functions}

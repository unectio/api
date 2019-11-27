package apilet

import (
	"github.com/unectio/util/restmux/client"
)

var FnEnvironment	= &client.Property{"env", Functions}
var FnLogs		= &client.Property{"logs", Functions}
var FnStats		= &client.Property{"stats", Functions}

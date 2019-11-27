package api

const (
	CalletPrefix		= "call"
	URLFunction		= rune('u')
	URLRouter		= rune('r')

	AuthCommonHeader	= "Authorization"
)

type ComputeImage struct {
	Cookie		string		`json:"cookie"`
}

type DepReqImage struct {
	Op		string		`json:"op"`
	Dep		*DepDescImage	`json:"dep"`
}

type DepDescImage struct {
	Proj		string		`json:"proj"`
	Lang		string		`json:"lang"`
	Class		string		`json:"class"`
}

func ParseCalletURL(val string) (rune, string) {
	return rune(val[0]), val[1:]
}

func MakeCalletFnURL(addr, cookie string) string {
	return addr + "/" + CalletPrefix + "/" + string(URLFunction) + cookie
}

func MakeCalletRouterURL(addr, cookie string) string {
	return addr + "/" + CalletPrefix + "/" + string(URLRouter) + cookie
}

func MakeCalletRunURL(fnid, ver string) string {
	return "/v1/run/" + fnid + "/" + ver
}

package api

type BuildRequest struct {
	Dir	string			`json:"dir"`
	PkgDir	string			`json:"pkgs"`
}

type BuildResponse struct {
	Code	int			`json:"code"`
	Stdout	string			`json:"stdout"`
	Stderr	string			`json:"stderr"`
}

type LangInfoResponse struct {
	Version	string			`json:"version"`
}

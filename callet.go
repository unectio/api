/////////////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2019-2020, Unectio Inc, All Right Reserved.
//
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
/////////////////////////////////////////////////////////////////////////////////

package api

const (
	CalletPrefix = "call"
	URLFunction  = rune('u')
	URLRouter    = rune('r')

	AuthCommonHeader = "Authorization"
)

type ComputeImage struct {
	Cookie string `json:"cookie"`
}

type DepReqImage struct {
	Op  string        `json:"op"`
	Dep *DepDescImage `json:"dep"`
}

type DepDescImage struct {
	Proj  string `json:"proj"`
	Lang  string `json:"lang"`
	Class string `json:"class"`
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

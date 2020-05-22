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

import (
	"time"
)

type LangImage struct {
	Language	string			`json:"language"`
	Build		bool			`json:"build"`

	*LangInfoResponse			`json:",inline,omitempty"`
}


type PkgImage struct {
	Id		ObjectId		`json:"id"`
	Name		string			`json:"name"`
	Version		string			`json:"version"`
}

type PkgStatsImage struct {
	DU		uint64			`json:"du_kbytes"`
}

func (ps *PkgStatsImage)SetDU(bytes uint64) {
	ps.DU = bytes >> 10
}

type MongoTierImage struct {
	Name		string			`json:"name"`
	Desc		string			`json:"desc"`
	Creds		string			`json:"creds,omitempty"`
}

type MwareImage struct {
	Type		string			`json:"type"`
}

type ProjectImage struct {
	Id		ObjectId		`json:"id"`
	Name		string			`json:"name"`
	Role		string			`json:"role"`
	UserData	string			`json:"userdata,omitempty"`
}

type FunctionImage struct {
	CommonImage				`json:",inline"                yaml:",inline"`

	State		string			`json:"state"                  yaml:"-"`

	Limits		*FuncLimitsImage	`json:"limits,omitempty"       yaml:"limits,omitempty"`
	Env		[]string		`json:"env,omitempty"          yaml:"env,omitempty"`
	CodeBalancer	string			`json:"code_balancer,omitempty yaml:"code_balancer,omitempty"`
}

type CodeImage struct {
	CommonImage				`json:",inline"                yaml:",inline"`

	Gen		int			`json:"generation"             yaml:"-"`
	State		string			`json:"state"                  yaml:"-"`
	Weight		int			`json:"weight"                 yaml:"-"`

	Lang		string			`json:"lang"                   yaml:"lang"`
	Source		*SourceImage		`json:"source,omitempty"       yaml:"source"`
}

type SourceImage struct {
	Text		[]byte			`json:"text_base64,omitempty"  yaml:"-"`
	URL		string			`json:"url,omitempty"          yaml:"url,omitempty"`
	RepoId		ObjectId		`json:"repo,omitempty"         yaml:"-"`
	Path		string			`json:"path,omitempty"         yaml:"path,omitempty"`

	Sync		bool			`json:"sync,omitempty"         yaml:"-"`
}

func (fc *SourceImage)RepoSource() bool {
	/* path can be empty, we'll fail opening it */
	return fc.RepoId != ""
}

type FuncLimitsImage struct {
	TmoMsec		*int			`json:"tmo_msec,omitempty"     yaml:"tmo_msec,omitempty"`
	Burst		*int			`json:"burst,omitempty"        yaml:"burst,omitempty"`
	Rate		*int			`json:"rate,omitempty"         yaml:"rate,omitempty"`
	Class		string			`json:"class,omitempty"        yaml:"class,omitempty"`
}

type FuncTargetImage struct {
	CommonImage				`json:",inline"                yaml:",inline"`
	Fn		*NextFunctionImage	`json:"function,omitempty"     yaml:"function,omitempty"`
	Success		ObjectId		`json:"on_success"             yaml:"on_success"`
	Failure		ObjectId		`json:"on_failure"             yaml:"on_failure"`
}

type NextFunctionImage struct {
	Id		ObjectId		`json:"id"                     yaml:"id"`
}

type FuncTriggerImage struct {
	CommonImage				`json:",inline"                yaml:",inline"`

	URL		*URLTrigImage		`json:"url,omitempty"          yaml:"utl,omitempty"`
	Cron		*CronTrigImage		`json:"cron,omitempty"         yaml:"cron,omitempty"`
	Websock		*WsTrigImage		`json:"websock,omitempty"      yaml:"websock,omitempty"`
	Event		*EventTrigImage		`json:"event,omitempty"        yaml:"event,omitempty"`
	CallKey		string			`json:"key"                    yaml:"key"`
	SortKey		string			`json:"sort,omitempty"         yaml:"sort,omitempty"`
}

func (ti *FuncTriggerImage)Src() string {
	switch {
	case ti.URL != nil:
		return "url"
	case ti.Cron != nil:
		return "cron"
	case ti.Websock != nil:
		return "websock"
	case ti.Event != nil:
		return "event"
	default:
		return ""
	}
}

type URLTrigImage struct {
	URL		string			`json:"url"                    yaml:"url"`
	AuthId		ObjectId		`json:"auth"                   yaml:"auth"`
}

type CronTrigImage struct {
	Tab		string			`json:"tab"                    yaml:"tab"`
	Args		map[string]string	`json:"args"                   yaml:"args"`
}

type WsTrigImage struct {
	WsId		ObjectId		`json:"websock"                yaml:"websock"`
}

type EventTrigImage struct {
	Source		ObjectId		`json:"source"                 yaml:"source"`
	/* XXX -- filters */
}

type FuncRun struct {
	Req		RunRequest		`json:"run"`
	Code		*SourceImage		`json:"code,omitempty"`
}

const CodeVersionLen = 16

type FuncWait struct {
	Tmo		uint64			`json:"tmo_msec"`
	Event		string			`json:"event"`
}

func (fw *FuncWait)Timeout() time.Duration {
	return time.Duration(fw.Tmo) * time.Millisecond
}

type FuncStatsImage struct {
	Calls		uint64			`json:"calls"`
	RunTime		uint64			`json:"runtime_us"`
	LastCall	string			`json:"lastcall"`
}

type ProjectStatsImage struct {
	Calls		uint64			`json:"calls"`
	RunTime		uint64			`json:"runtime_us"`
}

type LogEntry struct {
	Time		string			`json:"time"`
	Event		string			`json:"event"`
	Text		string			`json:"text"`
}

const (
	RepoLocal	string = "local"
)

type RepoImage struct {
	CommonImage				`json:",inline"`

	State		string			`json:"state"`

	Type		string			`json:"type"`
	URL		string			`json:"url,omitempty"`
	Head		string			`json:"head,omitempty"`

	Synced		string			`json:"synced_at"`
	SyncMinutes	*int			`json:"sync_minutes,omitempty"`
	Desc		bool			`json:"desc,omitempty"`
}

type RepoFileImage struct {
	Type		string			`json:"type"`
	Name		string			`json:"name,omitempty"`
	Path		string			`json:"path,omitempty"`
	Kids		*[]*RepoFileImage	`json:"kids,omitempty"`
}

type RepoSpec struct {
	Descr		string			`json:"desc" yaml:"desc"`
	Contents	[]*RepoEntry		`json:"files" yaml:"files"`
}

type RepoEntry struct {
	Name		string			`json:"name" yaml:"name"`
	Path		string			`json:"path" yaml:"path"`
	Desc		string			`json:"desc" yaml:"desc"`
	Lang		string			`json:"lang,omitempty" yaml:"lang,omitempty"`
}

type RouterImage struct {
	CommonImage				`json:",inline"                yaml:",inline"`

	AuthId		ObjectId		`json:"auth"                   yaml:"auth"`
	URL		string			`json:"url"                    yaml:"url"`

	Mux		[]*RouteRuleImage	`json:"mux,omitempty"          yaml:"mux,omitempty"`
}

type RouteRuleImage struct {
	Methods		string			`json:"methods"                yaml:"metods"` /* ,-separated */
	Path		string			`json:"path"                   yaml:"path"`
	Key		string			`json:"key"                    yaml:"key"`
	FnId		ObjectId		`json:"function"               yaml:"function"`
}

type AuthMethodImage struct {
	CommonImage				`json:",inline"                yaml:",inline"`

	JWT		*AuthJWTImage		`json:"jwt,omitempty"          yaml:"jwt,omitempty"`
	Platform	bool			`json:"platform,omitempty"     yaml:"platform,omitempty"`
}

type AuthJWTImage struct {
	Key		string			`json:"key"                    yaml:"key"`
}

type SecretImage struct {
	CommonImage				`json:",inline"`

	Payload		map[string]string	`json:"payload"`
	Reveal		string			`json:"reveal,omitempty"`
}

type WebsockImage struct {
	CommonImage				`json:",inline"                yaml:",inline"`

	AuthId		ObjectId		`json:"auth"                   yaml:"auth"`
	URL		string			`json:"url,omitempty"          yaml:"url"`
}

type MongoDbImage struct {
	CommonImage				`json:",inline"                yaml:",inline"`

	Tier		string			`json:"tier,omitempty"         yaml:"tier"`
}

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
	"encoding/json"
	"time"
)

const (
	EnvPrefix   = "LANGLET_"
	EnvPodIp    = EnvPrefix + "POD_IP"
	EnvPort     = EnvPrefix + "PORT"
	EnvInstance = EnvPrefix + "INSTANCE"
	EnvVolume   = EnvPrefix + "VOLUME"
	EnvProject  = EnvPrefix + "PROJECT"
	EnvLang     = EnvPrefix + "FN_LANG"
	EnvClass    = EnvPrefix + "CLASS"
	EnvToken    = EnvPrefix + "TOKEN"
)

const (
	VolSources  = "functions"
	VolRepos    = "repos"
	VolPackages = "packages"
	FnDescFile  = "desc.yaml"
)

const (
	RRCodeBalancer = "rr"
)

func DefaultCodeBalancer(cb string) string {
	if cb == "" {
		return RRCodeBalancer
	} else {
		return cb
	}
}

type FunctionDesc struct {
	Gen   int      `json:"gen"`
	TmoMs int      `json:"timeout_ms"`
	Env   []string `json:"env"`
}

func (d *FunctionDesc) Tmo() time.Duration {
	return time.Millisecond * time.Duration(d.TmoMs)
}

const (
	CallProjectHeader = "X-Target-ID"
	CallURLPrefix     = "/call/"
	CallGenParam      = "gen"

	MethodCron      = "CRON"
	MethodWebsocket = "WEBSOCK"
	MethodEvent     = "EVENT"
)

/*
 * The runner is responsible for setting
 * - Res      -- an encoded message from Main's 1st return
 * - Status   -- if provided by Main's 2nd return
 * - DeferMs  -- the same
 * - Chain    -- the same
 *
 * The watchdog is setting the rest
 * - Out, Err -- stdio streams
 * - LatUs    -- time taken
 */

type RunResponse struct {
	Status int             `json:"status"`
	Res    json.RawMessage `json:"res"`
	Out    string          `json:"out"`
	Err    string          `json:"err"`
	LatUs  uint            `json:"lat_us"`

	DeferMs uint     `json:"defer_ms,omitempty"`
	Next    *RunNext `json:"next,omitempty"`
}

func (resp *RunResponse) Defer() time.Duration {
	return time.Duration(resp.DeferMs) * time.Millisecond
}

func (resp *RunResponse) SetLat(since time.Time) {
	resp.LatUs = uint(time.Since(since) / time.Microsecond)
}

func (resp *RunResponse) GetLat() time.Duration {
	return time.Duration(resp.LatUs) * time.Microsecond
}

func CallURL(fnid, fgen, code, cgen string) string {
	/* http://langlet:port/fn_id/code_id?gen=fgen.cgen */
	return CallURLPrefix + fnid + "/" + code + "?" + CallGenParam + "=" + fgen + "." + cgen
}

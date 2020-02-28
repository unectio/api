package api

import (
	"time"
	"encoding/json"
)

const (
	EnvPrefix	= "LANGLET_"
	EnvPodIp	= EnvPrefix + "POD_IP"
	EnvPort		= EnvPrefix + "PORT"
	EnvInstance	= EnvPrefix + "INSTANCE"
	EnvVolume	= EnvPrefix + "VOLUME"
	EnvProject	= EnvPrefix + "PROJECT"
	EnvLang		= EnvPrefix + "FN_LANG"
	EnvClass	= EnvPrefix + "CLASS"
	EnvToken	= EnvPrefix + "TOKEN"
)

const (
	VolSources	= "functions"
	VolRepos	= "repos"
	VolPackages	= "packages"
	FnDescFile	= "desc.yaml"
)

type FunctionDesc struct {
	Gen		int		`json:"gen"`
	TmoMs		int		`json:"timeout_ms"`
	Env		[]string	`json:"env"`
}

func (d *FunctionDesc)Tmo() time.Duration {
	return time.Millisecond * time.Duration(d.TmoMs)
}

const (
	CallProjectHeader	= "X-Target-ID"
	CallURLPrefix		= "/call/"
	CallGenParam		= "gen"

	MethodCron		= "CRON"
	MethodWebsocket		= "WEBSOCK"
	MethodEvent		= "EVENT"
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
	Status	int			`json:"status"`
	Res	json.RawMessage		`json:"res"`
	Out	string			`json:"out"`
	Err	string			`json:"err"`
	LatUs	uint			`json:"lat_us"`

	DeferMs	uint			`json:"defer_ms,omitempty"`
	Chain	*ChainResponse		`json:"chain,omitempty"`
}

/*
 * Chain responce -- what the function wants us to call next
 */
type ChainResponse struct {
	Target	string			`json:"target"`
	Args	map[string]string	`json:"args"`
}

func (resp *RunResponse)Defer() time.Duration {
	return time.Duration(resp.DeferMs) * time.Millisecond
}

func (resp *RunResponse)SetLat(since time.Time) {
	resp.LatUs = uint(time.Since(since) / time.Microsecond)
}

func (resp *RunResponse)GetLat() time.Duration {
	return time.Duration(resp.LatUs) * time.Microsecond
}

func CallURL(fnid, fgen, code, cgen string) string {
	/* http://langlet:port/fn_id/code_id?gen=fgen.cgen */
	return CallURLPrefix + fnid + "/" + code + "?" + CallGenParam + "=" + fgen + "." + cgen
}

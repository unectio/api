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

type RunRequest struct {
	/*
	 * An arbitrary string. It's set by the user configuring a trigger
	 * (the .key field) or when configuring a router (route rule's 
	 * .key field). Use this string in the function's code to identify 
	 * specific trigger or router invoking the function (when the same 
	 * function is used by multiple triggers or routers)
	 */
	Key	string			`json:"key"             bson:"key"`
	/*
	 * url: http request method.
	 * other triggers: set to their "methods".
	 */
	Method	string			`json:"method"          bson:"method"`
	/*
	 * url: Part of the request URL after the function addreess itself.
	 * E.g. if your function trigger's URL is http://hoster.io/u12345
	 * and the function is called via http://hoster.io/u12345/foo/bar
	 * then the .Path will be /foo/bar. Same for router case.
	 *
	 * cron and websocket: empty
	 */
	Path	string			`json:"path"            bson:"path"`
	/*
 	 * url and websocket: JWT claims (only claims!) if the JWT auth
	 * method is set on the trigger or router and the caller has
	 * provided good token.
	 *
	 * If the auth method is "plain", then it will contain the
	 * "Authorization": <value> pair taken from headers.
	 *
	 * In the websocket case the authorization is done on the client
	 * connect-upgrade stage.
	 *
	 * cron: leaves it empty.
	 */
	Claims	map[string]interface{}	`json:"claims"          bson:"claims"`
	/*
	 * url: query args
	 * cron: trigger args
	 * websocket:
	 *	"websocket":    <ID of the websocket>
	 *	"conid":        <ID of the connection>
	 */
	Args	map[string]string	`json:"args"            bson:"args"`
	/*
	 * url: the content-type header
	 * websocket: message type
	 * cron: empty
	 */
	Content	string			`json:"content_type"    bson:"content_type"`
	/*
	 * url: the request body itself. If the .Content is application/json, then
	 * the body is additionally auto-unmarshalled.
	 *
	 * websocket: the message received
	 *
	 * cron: empty
	 */
	Body	[]byte			`json:"body"            bson:"body"`
}

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

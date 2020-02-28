package api

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


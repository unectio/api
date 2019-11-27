package api

const (
	ReqWebsockId string	= "websocket"
	ReqConnId string	= "conid"
)

type WebsockConnImage struct {
	Id		ObjectId	`json:"id"`
	From		string		`json:"from"`
	Subprot		string		`json:"subprotocol"`
}

type WebsockMsgImage struct {
	Type		int		`json:"type"`
	Message		[]byte		`json:"message"`
}

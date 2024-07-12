package SocketClient

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	websocketStringMessageType = 0
	websocketIntMessageType    = 1
	websocketBoolMessageType   = 2
	websocketJSONMessageType   = 4
	websocketMessagePrefix     = "gin-websocket-message:"
	websocketMessageSeparator  = ";"
	websocketMessagePrefixLen  = len(websocketMessagePrefix)
)

var (
	websocketMessageSeparatorLen    = len(websocketMessageSeparator)
	websocketMessagePrefixAndSepIdx = websocketMessagePrefixLen + websocketMessageSeparatorLen - 1
	websocketMessagePrefixIdx       = websocketMessagePrefixLen - 1
	websocketMessageSeparatorIdx    = websocketMessageSeparatorLen - 1
)

type (
	onConnectFunc                = func()
	onWebsocketDisconnectFunc    = func()
	onWebsocketNativeMessageFunc = func(websocketMessage string)
	onMessageFunc                = func(message any)
)

type SocketClient struct {
	conn                   *websocket.Conn
	isReady                bool
	connectListeners       []onConnectFunc
	disconnectListeners    []onWebsocketDisconnectFunc
	nativeMessageListeners []onWebsocketNativeMessageFunc
	messageListeners       map[string][]onMessageFunc
}

func NewSocketClient(endpoint string) *SocketClient {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dialer.Dial(endpoint, nil)
	if err != nil {
		log.Println(err)
	}
	var connectListeners []onConnectFunc
	var messageListeners map[string][]onMessageFunc
	var disconnectListeners []onWebsocketDisconnectFunc
	var nativeMessageListeners []onWebsocketNativeMessageFunc
	return &SocketClient{
		conn:                   conn,
		connectListeners:       connectListeners,
		messageListeners:       messageListeners,
		disconnectListeners:    disconnectListeners,
		nativeMessageListeners: nativeMessageListeners,
	}
}

func (s *SocketClient) isNumber(obj any) bool {
	return false
}

func (s *SocketClient) isString(obj any) bool {
	return false
}

func (s *SocketClient) isBoolean(obj any) bool {
	return false
}

func (s *SocketClient) isJSON(obj any) bool {
	return false
}

func (s *SocketClient) _msg(event string, websocketMessageType int, dataMessage string) string {
	return websocketMessagePrefix + event + websocketMessageSeparator + string(rune(websocketMessageType)) + websocketMessageSeparator + dataMessage
}

func (s *SocketClient) encodeMessage(event string, data any) string {
	return ""
}

func (s *SocketClient) decodeMessage(event string, websocketMessage string) any {
	return ""
}

func (s *SocketClient) getWebsocketCustomEvent(websocketMessage string) string {
	return ""
}

func (s *SocketClient) getCustomMessage(websocketMessage string) string {
	return ""
}

func (s *SocketClient) messageReceivedFromConn(evt map[string]any) {

}

func (s *SocketClient) OnConnect(fn onConnectFunc) {

}

func (s *SocketClient) fireConnect() {

}

func (s *SocketClient) OnDisconnect(fn onConnectFunc) {

}

func (s *SocketClient) fireDisconnect() {

}

func (s *SocketClient) OnMessage(cb onWebsocketNativeMessageFunc) {

}

func (s *SocketClient) fireNativeMessage(websocketMessage string) {

}

func (s *SocketClient) On(event string, cb onMessageFunc) {

}

func (s *SocketClient) fireMessage(event string, message any) {

}

func (s *SocketClient) Disconnect() {

}

func (s *SocketClient) EmitMessage(websocketMessage string) {

}

func (s *SocketClient) Emit(event string, data any) {

}

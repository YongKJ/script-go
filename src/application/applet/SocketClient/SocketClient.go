package SocketClient

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"strings"
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
	done                            = make(chan struct{})
	interrupt                       = make(chan os.Signal, 1)
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
	client := &SocketClient{
		conn:                   conn,
		connectListeners:       connectListeners,
		messageListeners:       messageListeners,
		disconnectListeners:    disconnectListeners,
		nativeMessageListeners: nativeMessageListeners,
	}
	client.initData(err)
	return client
}

func (s *SocketClient) initData(err error) {
	s.onOpen(err)
	s.onMessage()
	s.onClose()
}

func (s *SocketClient) onOpen(err error) {
	if err != nil {
		return
	}
	s.fireConnect()
	s.isReady = true
}

func (s *SocketClient) onMessage() {
	go func() {
		defer close(done)
		for {
			_, message, err := s.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
			s.messageReceivedFromConn(string(message))
		}
	}()
}

func (s *SocketClient) onClose() {
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-done:
			s.fireDisconnect()
			return
		case <-interrupt:
			log.Println("interrupt")
			return
		}
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
	if len(websocketMessage) < websocketMessagePrefixAndSepIdx {
		return ""
	}
	str := websocketMessage[websocketMessagePrefixAndSepIdx:]
	return str[0:strings.Index(str, websocketMessageSeparator)]
}

func (s *SocketClient) getCustomMessage(event string, websocketMessage string) string {
	eventIdx := strings.Index(websocketMessage, event+websocketMessageSeparator)
	return websocketMessage[eventIdx+len(event)+len(websocketMessageSeparator)+2:]
}

func (s *SocketClient) messageReceivedFromConn(message string) {
	if strings.Index(message, websocketMessageSeparator) != -1 {
		event := s.getWebsocketCustomEvent(message)
		if event != "" {
			s.fireMessage(event, s.getCustomMessage(event, message))
		}
	}
	s.fireNativeMessage(message)
}

func (s *SocketClient) OnConnect(fn onConnectFunc) {
	if s.isReady {
		fn()
	}
	s.connectListeners = append(s.connectListeners, fn)
}

func (s *SocketClient) fireConnect() {
	for i := 0; i < len(s.connectListeners); i++ {
		s.connectListeners[i]()
	}
}

func (s *SocketClient) OnDisconnect(fn onConnectFunc) {
	s.disconnectListeners = append(s.disconnectListeners, fn)
}

func (s *SocketClient) fireDisconnect() {
	for i := 0; i < len(s.disconnectListeners); i++ {
		s.disconnectListeners[i]()
	}
}

func (s *SocketClient) OnMessage(cb onWebsocketNativeMessageFunc) {
	s.nativeMessageListeners = append(s.nativeMessageListeners, cb)
}

func (s *SocketClient) fireNativeMessage(websocketMessage string) {
	for i := 0; i < len(s.nativeMessageListeners); i++ {
		s.nativeMessageListeners[i](websocketMessage)
	}
}

func (s *SocketClient) On(event string, cb onMessageFunc) {
	if _, ok := s.messageListeners[event]; !ok {
		var messageFunc []onMessageFunc
		s.messageListeners[event] = messageFunc
	}
	s.messageListeners[event] = append(s.messageListeners[event], cb)
}

func (s *SocketClient) fireMessage(event string, message any) {
	for key := range s.messageListeners {
		if key != event {
			continue
		}
		for i := 0; i < len(s.messageListeners[key]); i++ {
			s.messageListeners[key][i](message)
		}
	}
}

func (s *SocketClient) Disconnect() {
	err := s.conn.Close()
	if err != nil {
		log.Println(err)
	}
}

func (s *SocketClient) EmitMessage(websocketMessage string) {
	err := s.conn.WriteMessage(websocket.TextMessage, []byte(websocketMessage))
	if err != nil {
		log.Println(err)
	}
}

func (s *SocketClient) Emit(event string, data any) {
	messageStr := s.encodeMessage(event, data)
	s.EmitMessage(messageStr)
}

package SocketClient

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"reflect"
	"script-go/src/application/util/GenUtil"
	"strings"
	"time"
)

const (
	writeWait                  = 10 * time.Second
	pongWait                   = 60 * time.Second
	pingPeriod                 = (pongWait * 9) / 10
	maxMessageSize             = 512
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
	send                            = make(chan []byte, 256)
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
	client.initData()
	return client
}

func (s *SocketClient) initData() {
	s.onOpen()
	s.onMessage()
	s.onClose()
}

func (s *SocketClient) onOpen() {
	err := s.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		log.Println(err)
		return
	}
	if err = s.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
		return
	}
	s.isReady = true
	s.fireConnect()
	s.writePump()
}

func (s *SocketClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case message, ok := <-send:
				err := s.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err != nil {
					log.Println(err)
					return
				}
				if !ok {
					err = s.conn.WriteMessage(websocket.CloseMessage, []byte{})
					if err != nil {
						log.Println(err)
						return
					}
					return
				}

				w, err := s.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					log.Println(err)
					return
				}
				_, err = w.Write(message)
				if err != nil {
					log.Println(err)
					return
				}

				if err = w.Close(); err != nil {
					return
				}
			case <-ticker.C:
				err := s.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err != nil {
					log.Println(err)
					return
				}
				if err = s.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
			}
		}
	}()
}

func (s *SocketClient) onMessage() {
	s.conn.SetReadLimit(maxMessageSize)
	err := s.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		log.Println(err)
		return
	}
	s.conn.SetPongHandler(func(string) error {
		err = s.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
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

func (s *SocketClient) isFloat64(obj any) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Float64
}

func (s *SocketClient) isInt(obj any) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Int
}

func (s *SocketClient) isString(obj any) bool {
	return reflect.TypeOf(obj).Kind() == reflect.String
}

func (s *SocketClient) isBoolean(obj any) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Bool
}

func (s *SocketClient) isJSON(obj any) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Map
}

func (s *SocketClient) isNil(obj any) bool {
	return reflect.ValueOf(obj).IsNil()
}

func (s *SocketClient) _msg(event string, websocketMessageType int, dataMessage string) string {
	return websocketMessagePrefix + event + websocketMessageSeparator + string(rune(websocketMessageType)) + websocketMessageSeparator + dataMessage
}

func (s *SocketClient) encodeMessage(event string, data any) string {
	m := ""
	t := 0
	if s.isInt(data) {
		t = websocketIntMessageType
		m = GenUtil.IntToString(data.(int))
	} else if s.isFloat64(data) {
		t = websocketBoolMessageType
		m = GenUtil.Float64ToString(data.(float64))
	} else if s.isBoolean(data) {
		t = websocketBoolMessageType
		m = GenUtil.BoolToString(data.(bool))
	} else if s.isString(data) {
		t = websocketStringMessageType
		m = data.(string)
	} else if s.isJSON(data) {
		t = websocketJSONMessageType
		m = GenUtil.MapToString(data.(map[string]any))
	} else if !s.isNil(data) {
		log.Println("unsupported type of input argument passed, try to not include this argument to the 'Emit'")
	}
	return s._msg(event, t, m)
}

func (s *SocketClient) decodeMessage(event string, websocketMessage string) any {
	skipLen := websocketMessagePrefixLen + websocketMessageSeparatorLen + len(event) + 2
	if len(websocketMessage) < skipLen+1 {
		return nil
	}
	websocketMessageType := GenUtil.StrToInt(websocketMessage[skipLen-2:])
	theMessage := websocketMessage[skipLen:]
	if websocketMessageType == websocketIntMessageType {
		return GenUtil.StrToInt(theMessage)
	} else if websocketMessageType == websocketBoolMessageType {
		return GenUtil.StrToBoolean(theMessage)
	} else if websocketMessageType == websocketStringMessageType {
		return theMessage
	} else if websocketMessageType == websocketJSONMessageType {
		return GenUtil.StrToMap(theMessage)
	} else {
		return nil
	}
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

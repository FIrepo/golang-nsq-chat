package config

import "time"

const (
	// SocketBufferSize set buffer size for websocket connection
	SocketBufferSize = 1024

	// MessageBufferSize set buffer size for websocket message
	MessageBufferSize = 256

	// MaxMessageSize set max size allow for Websocket message
	MaxMessageSize = 512

	// PongWait set limit wait for receive messages from client
	PongWait = 30 * time.Second

	// PingPeriod set time to ping client
	PingPeriod = PongWait * 9 / 10

	// WriteWait set limit wait for writing to client
	WriteWait = 5 * time.Second
)

package routers

import "IM/server/websocket"

func WebsocketInit() {
	websocket.Register("ping", websocket.Ping)
	websocket.Register("login", websocket.Login)
	websocket.Register("heartbeat", websocket.Heartbeat)
}

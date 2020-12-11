package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"runtime/debug"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

type login struct {
	AppId  string
	Uid    string
	Client *Client
}

func (l *login) GetKey() string {
	return l.AppId + l.Uid
}

type Client struct {
	Addr          string
	Socket        *websocket.Conn
	Send          chan []byte
	AppId         string
	Uid           string
	FirstTime     int64
	HeartBeatTime int64
	LoginTime     int64
}

func NewClient(addr string, socket *websocket.Conn, firstTime int64) (client *Client) {
	client = &Client{
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartBeatTime: firstTime,
	}

	return
}

func (c *Client) GetKey() (key string) {
	key = c.AppId + c.Uid

	return
}

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println(message)
		ProcessData(c, message)
	}
}

func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		clientManager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				return
			}
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// 读取客户端数据
func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	c.Send <- msg
}

// 关闭
func (c *Client) close() {
	close(c.Send)
}

// 登录
func (c *Client) Login(appId string, uid string, loginTime int64) {
	c.AppId = appId
	c.Uid = uid
	c.LoginTime = loginTime
	c.HeartBeat(loginTime)
}

func (c *Client) HeartBeat(currentTime int64) {
	c.HeartBeatTime = currentTime

	return
}

func (c *Client) IsHeartbeatTimeout(currentTime int64) bool {
	return c.HeartBeatTime+heartbeatExpirationTime <= currentTime
}

func (c *Client) IsLogin() bool {
	return c.Uid != ""
}

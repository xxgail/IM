package websocket

import (
	"IM/lib/cache"
	"fmt"
	"sync"
	"time"
)

type ClientManager struct {
	Clients     map[*Client]bool // 所有连接的客户端
	ClientsLock sync.RWMutex
	Users       map[string]*Client // 所有登录的用户
	UsersLock   sync.RWMutex
	Register    chan *Client // 连接事件
	Login       chan *login  // 登录事件
	Unregister  chan *Client // 退出登录
	Broadcast   chan []byte  // 广播 向全部成员发送数据
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Login:      make(chan *login, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}

	return
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.EventRegister(conn)
		case login := <-manager.Login:
			manager.EventLogin(login)
		case conn := <-manager.Unregister:
			manager.EventUnregister(conn)
		}
	}
}

func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	_, ok = manager.Clients[client]

	return
}

// 获取所有
func (manager *ClientManager) GetClients() (clients map[*Client]bool) {

	clients = make(map[*Client]bool)

	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value

		return true
	})

	return
}

// 遍历
func (manager *ClientManager) ClientsRange(f func(client *Client, value bool) (result bool)) {

	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	for key, value := range manager.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}

	return
}

func (manager *ClientManager) GetClientsLen() (clientsLen int) {
	clientsLen = len(manager.Clients)
	return
}

// 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	manager.Clients[client] = true
}

// 删除客户端
func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	if _, ok := manager.Clients[client]; ok {
		delete(manager.Clients, client)
	}
}

// 获取用户的连接
func (manager *ClientManager) GetUserClient(appId string, userId string) (client *Client) {
	manager.UsersLock.RLock()
	defer manager.UsersLock.RUnlock()

	userKey := appId + userId
	if value, ok := manager.Users[userKey]; ok {
		client = value
	}

	return
}

// GetClientsLen
func (manager *ClientManager) GetUsersLen() (userLen int) {
	userLen = len(manager.Users)

	return
}

// 添加用户
func (manager *ClientManager) AddUsers(key string, client *Client) {
	manager.UsersLock.Lock()
	defer manager.UsersLock.Unlock()

	manager.Users[key] = client
}

// 删除用户
func (manager *ClientManager) DelUsers(client *Client) (result bool) {
	manager.UsersLock.Lock()
	defer manager.UsersLock.Unlock()

	key := client.AppId + client.Uid
	if value, ok := manager.Users[key]; ok {
		// 判断是否为相同的用户
		if value.Addr != client.Addr {
			return
		}
		delete(manager.Users, key)
		result = true
	}

	return
}

// 建立连接
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)
}

// 登录
func (manager *ClientManager) EventLogin(login *login) {
	client := login.Client

	if manager.InClient(client) {
		userKey := login.GetKey()
		manager.AddUsers(userKey, client)
	}
}

// 用户断开连接
func (manager *ClientManager) EventUnregister(client *Client) {
	manager.DelClients(client)

	// 删除用户连接
	deleteResult := manager.DelUsers(client)
	if deleteResult == false {
		return
	}

	// 清除redis登录数据
	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err == nil {
		userOnline.LogOut()
		_ = cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	}
}

// 获取用户所在的连接
func GetUserClient(appId string, userId string) (client *Client) {
	client = clientManager.GetUserClient(appId, userId)

	return
}

// 定时清理超时连接
func ClearTimeoutConnections() {
	currentTime := time.Now().Unix()

	clients := clientManager.GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			fmt.Println("心跳时间超时 关闭连接", client.Addr, client.Uid, client.LoginTime, client.HeartBeatTime)

			client.Socket.Close()
		}
	}
}

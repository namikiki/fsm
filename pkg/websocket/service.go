package websocket

//
//import (
//	"log"
//
//	"fsm/pkg/domain"
//
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//)
//
//type Client struct {
//	uid  string
//	cid  string
//	conn *websocket.Conn
//}
//
//func (c *Client) readPump(hub *SocketPool) {
//	defer func() {
//		hub.unregister <- c
//		c.conn.Close()
//	}()
//
//	for {
//		_, message, err := c.conn.ReadMessage()
//
//		if err != nil {
//			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
//				log.Println("客户端断开连接", err)
//			}
//			break
//		}
//
//		log.Println(string(message))
//	}
//}
//
//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}
//
//type SocketPool struct {
//	account    map[string]map[string]*websocket.Conn
//	register   chan *Client
//	unregister chan *Client
//}
//
//func NewService() domain.WebSocketService {
//	return &SocketPool{
//		account:    make(map[string]map[string]*websocket.Conn),
//		register:   make(chan *Client),
//		unregister: make(chan *Client),
//	}
//}
//
//func (p *SocketPool) Info() map[string]map[string]*websocket.Conn {
//	return p.account
//}
//
//func (p *SocketPool) Connect(c *gin.Context) error {
//	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		return err
//	}
//
//	client := &Client{
//		uid:  c.Query("uid"),
//		cid:  c.Query("cid"),
//		conn: conn,
//	}
//
//	if client.uid == "" || client.cid == "" {
//		conn.Close()
//		return nil
//	}
//
//	p.register <- client
//
//	//go client.writePump()
//	client.readPump(p)
//
//	return err
//}
//
//func (p *SocketPool) Close() error {
//	log.Println("ws Close")
//	return nil
//}
//
//func (p *SocketPool) Run() {
//	for {
//		select {
//		case client := <-p.register:
//			clients, ok := p.account[client.uid]
//			if ok {
//				clients[client.cid] = client.conn
//			} else {
//				clients = make(map[string]*websocket.Conn)
//				clients[client.cid] = client.conn
//				p.account[client.uid] = clients
//			}
//			log.Println("Client connected:", p.account[client.uid])
//
//		case client := <-p.unregister:
//			delete(p.account[client.uid], client.cid)
//			log.Println("删除客户端")
//		}
//	}
//}
//
//func (p *SocketPool) BroadcastFileDown(uid string, clientUUID string) {
//
//	for i, con := range p.account[uid] {
//		message := []byte(i + "back")
//		if err := con.WriteMessage(websocket.TextMessage, message); err != nil {
//			log.Println(err)
//			continue
//		}
//	}
//
//}

package msgServer

import (
	"log"
	"net"
	"encoding/json"
	"LemonChat/clients"
	"LemonChat/commons"
	"LemonChat/protocols"
)

type MessageServer struct {
	ip            string
	port          string
	lAddr         string
	listener      net.Listener
	clientManager *ClientManager
}

func NewMessageServer(ip string, port string, laddr string) (ms *MessageServer) {
	return &MessageServer{ip: ip, port: port, lAddr: laddr, clientManager: NewClientManager()}
}

func (ms *MessageServer) Serve() {
	Log("start listen!")
	listener, err := net.Listen("tcp", ms.lAddr)
	if err != nil {
		Log("create listener fail : ", err.Error())
		return
	}
	Log("listen success!")
	for {
		conn, err := listener.Accept()
		if err != nil {
			Log("client connect failed : ", err.Error())
			break
		}
		Log("accept successed!")

		go handleRead(conn, ms)
	}
}

func handleRead(conn net.Conn, ms *MessageServer) {
	tempBuffer := make([]byte, 0)
	buffer := make([]byte, 1024)
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel, ms, conn)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			continue
		}
		Log("Read =", string(buffer[:n]), " n = ", n)
		tempBuffer = protocols.Depack(append(tempBuffer, buffer[:n]...), readerChannel)
	}
}

func reader(readChannel chan []byte, ms *MessageServer, conn net.Conn) {
	for {
		select {
		case data := <-readChannel:
			var r request
			Log("received : ", string(data))
			err1 := json.Unmarshal(data, &r)
			if err1 != nil {
				Log("unmarshal failed : ", err1.Error())
				continue
			}
			handleRequest(r, ms, conn)
		}
	}
}

func handleRequest(r request, ms *MessageServer, conn net.Conn) {
	switch r.RequestHandle {
	case commons.LOGIN_REQUEST:
		Log("login request : ", r.ToString())
		c := &clients.Client{Id: r.ClientId, Conn: conn}
		ms.clientManager.AddClient(c)
		Log("client count : ", ms.clientManager.GetClientCount())
	case commons.SEND_MSG:
		Log("send msg:", r.ToString())
		sendMsg("Back Message!", conn)
	default:

	}

}

func sendMsg(msg string, conn net.Conn) {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		Log("Send fail!")
	}
}

func Log(v ...interface{}) {
	log.Println(v)
}

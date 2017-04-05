package clients

import "net"

type Client struct {
	Id string
	Conn net.Conn
	Person *Person
	LastHeartBeatTime int64
}

func NewClient(id string, conn net.Conn) (c *Client){
	c = &Client{Id: id, Conn: conn}
	return
}

func (c *Client) SetPerson() {

}

func (c *Client) SetLastHeartBeatTime(t int64) {

}

func (c *Client) Quit() {
	c.Conn.Close()
}

func (c *Client) GetId() string{
	return c.Id
}

func (c *Client) GetConn() net.Conn{
	return c.Conn
}

func (c *Client) GetPerson() *Person {
	return c.Person
}

func (c *Client) GetHeratBeatTime() int64 {
	return c.LastHeartBeatTime
}

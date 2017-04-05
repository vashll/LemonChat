package msgServer

import (
	"LemonChat/clients"
	"errors"
)

type ClientManager struct {
	clientMap map[string]*clients.Client
}

func NewClientManager() *ClientManager{
	cMap := make(map[string]*clients.Client)
	return &ClientManager{clientMap: cMap}
}

func (cm *ClientManager) AddClient(c *clients.Client) {
	cm.clientMap[c.GetId()] = c
}

func (cm *ClientManager) RemoveClient(clientId string) {
	_, exist := cm.clientMap[clientId]
	if exist {
		delete(cm.clientMap, clientId)
	}
}

func (cm *ClientManager) GetClient(id string) (*clients.Client, error) {
	c, exist := cm.clientMap[id]
	if exist {
		return c, nil
	}
	str := "Client : " + id + " not exist"
	return nil, errors.New(str)
}

func (cm *ClientManager) GetClientCount () int{
	return len(cm.clientMap)
}
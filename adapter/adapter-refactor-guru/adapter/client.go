package adapter

import "fmt"

type client struct {
}

func NewClient() *client {
	return &client{}
}

func (c *client) InsertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

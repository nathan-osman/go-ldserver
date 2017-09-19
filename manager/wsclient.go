package manager

import (
	"github.com/gorilla/websocket"
)

type wsState struct {
	Name  string `json:"name"`
	State bool   `json:"state"`
}

type wsClient struct {
	refCount  int
	conn      *websocket.Conn
	stateCh   chan *wsState
	stoppedCh chan bool
}

func (c *wsClient) run() {
	defer close(c.stoppedCh)
	defer c.conn.Close()
	for m := range c.stateCh {
		if err := c.conn.WriteJSON(m); err != nil {
			// TODO: better error handling
			continue
		}
	}
}

func newWSClient(host string) (*wsClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		return nil, err
	}
	c := &wsClient{
		conn:      conn,
		stateCh:   make(chan *wsState),
		stoppedCh: make(chan bool),
	}
	go c.run()
	return c, nil
}

func (c *wsClient) Add() {
	c.refCount++
}

func (c *wsClient) Send(s *wsState) {
	c.stateCh <- s
}

func (c *wsClient) Release() {
	if c.refCount--; c.refCount == 0 {
		c.Close()
	}
}

func (c *wsClient) Close() {
	close(c.stateCh)
	<-c.stoppedCh
}

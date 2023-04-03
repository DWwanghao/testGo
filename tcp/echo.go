package tcp

import (
	"bufio"
	"context"
	"io"
	"net"
	"pero-redis/lib/logger"
	"pero-redis/lib/sync/atomic"
	"pero-redis/lib/sync/wait"
	"sync"
	"time"
)

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

func MakeHandler() *EchoHandler {
	return &EchoHandler{}
}

func (e *EchoClient) Close() error {
	e.Waiting.WaitWithTimeout(10 * time.Second)
	_ = e.Conn.Close()
	return nil
}

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

func (e *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if e.closing.Get() {
		_ = conn.Close()
	}
	client := &EchoClient{
		Conn: conn,
	}
	e.activeConn.Store(client, struct{}{})
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				logger.Info("connection close")
				e.activeConn.Delete(client)
			} else {
				logger.Warn(err)
			}
			return
		}

		client.Waiting.Add(1)
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done()
	}
}

func (e *EchoHandler) Close() error {
	logger.Info("handler shutting down")
	e.closing.Set(true)
	e.activeConn.Range(func(key, value interface{}) bool {
		client := key.(*EchoClient)
		_ = client.Conn.Close()
		return true
	})
	return nil
}

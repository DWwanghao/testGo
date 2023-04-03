package tcp

import (
	"context"
	"net"
	"os"
	"os/signal"
	"pero-redis/interface/tcp"
	"pero-redis/lib/logger"
	"sync"
	"syscall"
)

type Config struct {
	Address string
}

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	closeChan := make(chan struct{})
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigChan
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	listen, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info("start listen")
	ListenAndServe(listen, handler, closeChan)
	return nil
}

func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) {
	go func() {
		<-closeChan
		logger.Info("shutting down")
		_ = listener.Close()
		_ = handler.Close()
	}()
	defer func() {
		_ = listener.Close()
		_ = handler.Close()
	}()
	background := context.Background()
	var waitDone sync.WaitGroup
	for true {
		accept, err := listener.Accept()
		if err != nil {
			break
		}
		logger.Info("accepted link")
		waitDone.Add(1)
		go func() {
			defer func() {
				waitDone.Done()
			}()
			handler.Handle(background, accept)
		}()

	}
	waitDone.Wait()

}

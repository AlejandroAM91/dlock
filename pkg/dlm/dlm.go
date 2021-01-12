package dlm

import (
	"fmt"
    "net"
)

const (
	connType = "tcp"
)

type DLM struct {
	Peer

	l net.Listener
}

func NewDLM(host string, port string) *DLM {
	dlm := DLM{
		Peer: Peer{
			Host: host,
			Port: port,
		},
	}
	return &dlm
}

func (s *DLM) Start() error {
	var err error
	if s.l, err = net.Listen(connType, s.Peer.Host + ":" + s.Peer.Port); err != nil {
		return err
	}

	// TODO: Check this
	go s.accept()

	return nil
}

func (s DLM) Stop() error {
	if err := s.l.Close(); err != nil {
		return err
	}

	return nil
}

func (s DLM) accept() {
	// TODO: Check this
	for {
        c, err := s.l.Accept()
        if err != nil {
            fmt.Println("Error connecting")
            return
        }
        fmt.Println("Client connected.")

        fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

        // go handleConnection(c)
    }
}

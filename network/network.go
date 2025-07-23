package network

import (
	"bufio"
	"kv-database/compute"
	"log"
	"net"
)

type Network interface {
	StartListening() error
}

type networkImpl struct {
	cmp compute.Compute
}

func New() Network {
	return &networkImpl{
		cmp: compute.New(),
	}
}

func (n *networkImpl) StartListening() error {
	listener, err := net.Listen("tcp", ":9099")

	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		log.Printf("Connected user: %s\n", conn.RemoteAddr().String())

		go n.handleConnection(conn)
	}
}

func (n *networkImpl) handleConnection(conn net.Conn) {
	sc := bufio.NewScanner(conn)

	for sc.Scan() {
		if sc.Text() == "exit" {
			conn.Close()
			log.Printf("Disconnected user: %s\n", conn.RemoteAddr().String())
		}

		r, err := n.cmp.Exec(sc.Text())
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			if r.String() != "" {
				conn.Write([]byte(r.String() + "\n"))
			}
		}
	}
}

func (n *networkImpl) stopListening() {

}

package podcore

import (
	"fmt"
	"net"
	"sync"
)

type Transport string

const (
	TCP Transport = "tcp"
	UDP           = "udp"
)

type Server struct {
	IPAddr string
	Ports  []string
	Transp Transport
}

func NewServer(ipaddr string, ports []string, transport Transport) *Server {
	return &Server{IPAddr: ipaddr, Ports: ports, Transp: transport}
}

func (svr *Server) Start() {
	fmt.Println("server start", svr)
	var wg sync.WaitGroup
	wg.Add(len(svr.Ports))

	for _, port := range svr.Ports {
		go func(port string) {
			fmt.Println("servering on ", port)
			defer wg.Done()

			switch svr.Transp {
			case TCP:
				listen, err := net.Listen("tcp", ":"+port)
				if err != nil {
					fmt.Println(err)
					return
				}
				for {
					conn, err := listen.Accept()
					if err != nil {
						fmt.Println(err)
						return
					}
					go handleConnection(conn)
				}
			case UDP:
				pc, err := net.ListenPacket(string(svr.Transp), svr.IPAddr)
				if err != nil {
					return
				}
				buffer := make([]byte, 1024)
				for {
					n, addr, err := pc.ReadFrom(buffer)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Printf("packet received : bytes=%d from=%s buffer=%v", n, addr, buffer[:n])
				}
			}
		}(port)
	}
	wg.Wait()
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Printf("packet received : bytes=%d from=%s buffer=%v", n, conn.RemoteAddr(), buffer[:n])
}

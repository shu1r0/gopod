package main

import podserver "GoPod/podcore"

func main() {
	tcpSvr := podserver.NewServer("127.0.0.1", []string{"8888", "9999"}, podserver.TCP)
	tcpSvr.Start()
}

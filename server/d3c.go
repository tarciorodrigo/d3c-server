package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Server em execução")
	startListener("9090")
}

func startListener(port string) {
	// tcp->0.0.0.0:9090
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)

	if err != nil {
		log.Fatal("Erro ao iniciar o listener: ", err.Error())
	} else {
		for {
			canal, err := listener.Accept()
			defer canal.Close()

			if err != nil {
				log.Println("Erro em um novo canal:", err.Error())
			}

			log.Println("Nova conexão: ", canal.RemoteAddr().String())
		}
	}
}

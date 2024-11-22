package main

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"log"
	"net"
	"os"
	"time"

	"d3c/commons"
)

var (
	mensagem    commons.Mensagem
	tempoEspera = 10
)

const (
	SERVIDOR = "127.0.0.1"
	PORTA    = "9090"
)

func init() {
	mensagem.AgentHostName, _ = os.Hostname()
	mensagem.AgentCWS, _ = os.Getwd()
	mensagem.AgentID = geraID()
}

func main() {
	log.Println("Agente em execução")

	for {
		canal, err := conectaServidor()

		if err != nil {
			log.Println("Error : ", err.Error())
		} else {
			defer canal.Close()

			gob.NewEncoder(canal).Encode(mensagem)
			gob.NewDecoder(canal).Decode(mensagem)
		}

		time.Sleep(time.Duration(tempoEspera) * time.Second)
	}
}

func conectaServidor() (canal net.Conn, err error) {
	canal, err = net.Dial("tcp", SERVIDOR+":"+PORTA)

	return canal, err
}

func geraID() string {
	myTime := time.Now().String()

	hasher := md5.New()

	hasher.Write([]byte(mensagem.AgentHostName + myTime))

	return hex.EncodeToString(hasher.Sum(nil))
}

package server

import (
	packet "EGTS_PROTOCOL/pkg/egts/packet"
	"EGTS_PROTOCOL/pkg/egts/subrecord"
	"fmt"
	"log"
	"net"
)

func Handler(){

	listener, err := net.Listen("tcp", ":1924" )

	if err!=nil{
		log.Println("Error to listen server ", err)

	}
	defer listener.Close()

	log.Println("Server listening")

	for{
		conn, err := listener.Accept()

		if err!=nil{
			log.Println("Error to connection ", err)
			conn.Close()
			continue
		}

		go handlerConnection(conn)
	}

}


func handlerConnection(conn net.Conn){
	defer conn.Close()
	buff:=make([]byte, 65535)
	for{
		req, err :=conn.Read(buff)

		if err!=nil{
			log.Println("Read error ",err)
			return
		}

		readPacket, err:= packet.ReadPacket(buff[:req])
		if err!=nil{
			log.Println("Error to read packet ", err)
			return
		}

		data := *readPacket.ServicesFrameData.(*packet.ServicesFrameData)
		recordData := data[0].RecordsData[0]

		response := readPacket.PrepareAnswer(data[0].RecordNumber, readPacket.PacketID)

		switch recordData.SubrecordType {
			case packet.PosData:
				fmt.Println(recordData.SubrecordData.(*subrecord.SRPosData))
				break

		}

		_, err = conn.Write(response.Encode())
		if err != nil {
			log.Println("Error to write packet ", err)
			return
		}

	}

}







package conn

import (
	"encoding/binary"
	"gonat/proto"
	"net"
	"strconv"
)

func Start(port string) {

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	for {
		local_con, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go local_task(local_con)
	}

}

func local_task(local_con net.Conn) {

	port_b := make([]byte, 4, 4)
	_, err := local_con.Read(port_b)
	port := binary.BigEndian.Uint32(port_b)

	listen, err := net.Listen("tcp", ":"+strconv.Itoa(int(port)))
	if err != nil {
		local_con.Close()
		return
	}

	addr := listen.Addr().String()

	p := proto.Proto{proto.TCP_SEND_PROTO, 0, []byte(addr)}
	_, err = local_con.Write(p.Marshal())
	if err != nil {
		local_con.Close()
		return
	}

	start_conversation(listen, local_con)

}
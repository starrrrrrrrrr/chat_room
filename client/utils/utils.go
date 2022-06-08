package utils

import (
	"encoding/binary"
	"encoding/json"
	"exercise/chatroom/common/message"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (tf *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取数据...")
	_, err = tf.Conn.Read(tf.Buf[0:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	pkgLen := binary.BigEndian.Uint32(tf.Buf[0:4])
	n, err := tf.Conn.Read(tf.Buf[0:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	err = json.Unmarshal(tf.Buf[0:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

func (tf *Transfer) WritePkg(data []byte) (err error) {
	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(tf.Buf[:4], pkgLen)
	n, err := tf.Conn.Write(tf.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	n, err = tf.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	return
}

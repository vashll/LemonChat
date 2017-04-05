package protocols

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	ConstHeader       = "Header"
	ConstHeaderLength = len(ConstHeader)
	ConstMLength      = 4
)

func Enpack(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

func Depack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	protoConstLegth := ConstHeaderLength + ConstMLength
	for i = 0; i < length; i++ {
		//Log("Depack length = ", length, "i = ", i)
		//Log("Depack data ", string(buffer))
		if length < protoConstLegth {
			break
		}
		if i+protoConstLegth >= length {
			return make([]byte, 0)
		}
		Log("Depack length = ", length, "i = ", i)
		if string(buffer[i: i+ConstHeaderLength]) == ConstHeader {
			dataLength := BytesToInt(buffer[i+ConstHeaderLength: i+protoConstLegth])
			if length < i+protoConstLegth+dataLength {
				break
			}
			data := buffer[i+protoConstLegth: i+protoConstLegth+dataLength]
			readerChannel <- data
			i += protoConstLegth + dataLength - 1
		}

	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

func IntToBytes(len int) []byte {
	x := int32(len)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func Log(v ...interface{}) {
	log.Println(v)
}

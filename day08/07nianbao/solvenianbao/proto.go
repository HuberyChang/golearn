package solvenianbao

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 编码
func Encode(msg string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(msg))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) //读取前四个字节
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲区中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
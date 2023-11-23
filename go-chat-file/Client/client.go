package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"bufio"
	"strings"
	"path/filepath"
)

func sendFile(conn net.Conn, filePath string)  {
	defer conn.Close()

	//เปิดไฟล์ที่จะส่ง
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() 

	_, fileName := filepath.Split(filePath)

	conn.Write([]byte(fileName))

	buffer := make([]byte, 1024)
	for {
		//อ่านไฟล์ไปเก็บที่ buffer
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Send file success")
			}else {
				fmt.Println(err)
			}
			return
		}
		conn.Write(buffer[:n])
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Connect to server success")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter file path+name: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSuffix(filePath, "\n")

	sendFile(conn, filePath)
}
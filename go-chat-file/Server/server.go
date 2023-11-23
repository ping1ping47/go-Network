package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn)  {
	defer conn.Close()

	//สร้างตัวเก็บข้อมูล
	buffer := make([]byte, 1024)

	//รับไฟล์จาก Client
	fileNameBuffer:= make([]byte, 64)

	n, err := conn.Read(fileNameBuffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := string(fileNameBuffer[:n])
	fmt.Println("Receive File Name:", fileName)

	//สร้างไฟล์เก็บข้อมูล
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//รับข้อมูลจาก buffer มาใส่
	for{
		n, err := conn.Read(buffer)
		if err != nil{
			if err == io.EOF {
				fmt.Println("Transfer Complete")
			} else{
				fmt.Println(err)
			}
			return
		}
		//เขียนข้อมูลลงไฟล์
		file.Write(buffer[:n])
	}
}

func main(){
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 5000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Client Connected:", conn.RemoteAddr())

		go handleConnection(conn)
	}
}
package main

func main() {
	// 创建一个server句柄
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
// 创建服务器教程
package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}

// 创建并开启一个服务器

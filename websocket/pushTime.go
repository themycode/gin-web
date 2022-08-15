package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/getTime", SendTime)
	fmt.Println("已经注册事件")

	err := http.ListenAndServe(":13000", nil)
	if err != nil {
		fmt.Println("监听失败", err)
	}
}

func SendTime(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级失败", err)
		return
	}
	defer c.Close()

	// 接收消息
	mt, message, err := c.ReadMessage()
	if err != nil {
		fmt.Println("接收消息失败", err)
	} else {
		fmt.Println("接收到消息", string(message))
		var sendbuff string = "hello client!"
		if err = c.WriteMessage(mt, []byte(sendbuff)); err != nil {
			fmt.Println("发送消息失败", err)
		}
		fmt.Println("开始发送时间")
		for {
			sendbuff = time.Now().Format("2006-01-02 15:04:05")
			err = c.WriteMessage(mt, []byte(sendbuff))
			if err != nil {
				fmt.Println("发送时间错误", err)
				break
			}
			fmt.Println("发送时间", sendbuff)
			time.Sleep(5 * time.Second)
		}
		fmt.Println("停止发送时间")
	}
}

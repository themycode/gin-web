package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024 * 10,
		WriteBufferSize: 1024 * 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var memeberList map[string]*websocket.Conn

func salir() {
	for {
		for k, v := range memeberList {
			if err := v.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
				v.Close()
				delete(memeberList, k)
				fmt.Println("离开了", k)
				fmt.Println("当前连接数", len(memeberList))
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wbsCon *websocket.Conn
		err    error
		data   []byte
	)
	wbsCon, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级ws出错", err)
		return
	} else {
		memeberList[wbsCon.RemoteAddr().String()] = wbsCon
		fmt.Println("来了新连接", wbsCon.RemoteAddr().String())
		fmt.Println("当前连接数：", len(memeberList))
	}
	for {

		if _, data, err = wbsCon.ReadMessage(); err != nil {
			goto ERR
		}
		fmt.Println(wbsCon.RemoteAddr().String(), "说：", string(data))
		if err = wbsCon.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	wbsCon.Close()
}

func main() {
	fmt.Println("开始运行")
	memeberList = make(map[string]*websocket.Conn)
	// go salir()
	http.HandleFunc("/ws", wsHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package server

import (
	checkinMsg "checkin/server/msg"
	"encoding/json"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func handleGetuserlist(conn *websocket.Conn, stn bool) {
	sendData(conn, checkinMsg.GetuserlistMessage{
		Cmd: CmdGetuserlist,
		Stn: stn,
	})
}

func receiveGetuserlist(conn *websocket.Conn, msg []byte) {
	var response checkinMsg.GetuserlistResponse
	if err := json.Unmarshal(msg, &response); err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		return
	}
	if !response.Result {
		return
	}

	if response.Count == 0 {
		log.Println("获取用户列表结束")
		return
	}

	handleGetuserlist(conn, false)

}

func handleGetuserinfo(conn *websocket.Conn, msg checkinMsg.GetuserinfoMessage) {
	sendData(conn, msg)
}

func receiveGetuserinfo(conn *websocket.Conn, msg []byte) {

}

// HandleSetUserInfoAll 向所有已连接设备下发
func handleSetUserInfoAll(msg checkinMsg.SetuserinfoMessage) {

	for client := range clients {
		if err := client.WriteJSON(msg); err != nil {
			log.Println("Error sending message:", err)
			client.Close()
			delete(clients, client) // 移除失效的连接
		}
	}
}

// handleSetuserinfo设置
func handleSetuserinfo(conn *websocket.Conn, msg checkinMsg.SetuserinfoMessage) {
	if err := conn.WriteJSON(msg); err != nil {
		log.Println("Error sending message:", err)
		conn.Close()
		delete(clients, conn) // 移除失效的连接
	}
}

func receiveSetuserinfo(conn *websocket.Conn, msg []byte) {
	_ = conn
	var response checkinMsg.WSResponse
	if err := json.Unmarshal(msg, &response); err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		return
	}
	if !response.Result {
		log.Println("Error set user info:", response.Msg)
	} else {
		log.Printf("对设备[%s]下发用户信息[%d]成功", response.Sn, response.Enrollid)
	}

}

package main

type ChatMessage struct {
	Id       int    `json:"id"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Text     string `json:"text"`
	SendTime string `json:"sendTime"`
}

package data

type TcpSSetting struct {
	AddressText      string // 地址
	PortText         string // 端口
	ListenBut        bool
	IsHeartBeat      bool     // 是否启动心跳
	HeartFps         string   // 心跳帧率
	IsSendHello      bool     //是否发送hello
	CharacterSet     string   // 字符集编码
	IsAutomaticReply bool     // 是否自动回复
	LinkList         []string // 连接客户端列表
	SendData         string   // 发送的数据
	SendMsgList      []string // 已经发送的数据列表
	SendNums         int      // 已经发送的数据数量
	SendCharNums     int      // 已经发送的数据字符数量
	AcceptData       string   // 已经接受的数据
	AcceptMsgList    []string // 已经接受到的数据的列表
	AcceptNums       int      // 已经接受到的数据的数目
	AcceptCharNums   int
}

var TcpSSet TcpSSetting

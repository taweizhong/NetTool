package handle

import (
	"NetTool/data"
	"fyne.io/fyne/v2"
	"log"
	"net"
	"strings"
)

var TcpClientConn net.Conn

func TcpClient(address string, Ch chan net.Conn) {
	var err error
	TcpClientConn, err = net.Dial("tcp", address)
	if err != nil {
		log.Println("拨号失败")
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "拨号失败",
			Content: "Content:" + err.Error(),
		})
		return
	} else {
		log.Println("拨号成功")
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "拨号成功",
			Content: "",
		})
	}
	go bcc(TcpClientConn)
	Ch <- TcpClientConn
}
func bcc(conn net.Conn) {
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			return
		}
		str := string(buf[:n])
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "收到服务端的消息",
			Content: "Content:" + str,
		})
		data.TcpCSet.AcceptMsgList = append(data.TcpCSet.AcceptMsgList, str+"\n")
	}
}
func TcpClientSendData(data string) {
	data = strings.TrimSpace(data)
	_, err := TcpClientConn.Write([]byte(data))
	if err != nil {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "写入失败",
			Content: "Content:" + err.Error(),
		})
	}
}
func TcpClientClose(Ch chan net.Conn) {
	tcpClientConn := <-Ch
	err := tcpClientConn.Close()
	if err != nil {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Client关闭失败",
			Content: "Content:" + err.Error(),
		})
	} else {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Client关闭成功",
			Content: "",
		})
	}
}

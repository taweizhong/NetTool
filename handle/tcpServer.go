package handle

import (
	"NetTool/data"
	"fyne.io/fyne/v2"
	"net"
	"strings"
)

var TcpServerConn net.Conn

func TcpServer(address string, Ch chan net.Conn) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "监听失败",
			Content: "Content:" + err.Error(),
		})
		return
	} else {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "监听成功",
			Content: "",
		})

	}
	for {
		TcpServerConn, _ = listen.Accept()
		go acc(TcpServerConn)
		Ch <- TcpServerConn
	}
}

func acc(conn net.Conn) {
	data.TcpSSet.LinkList = append(data.TcpSSet.LinkList, conn.RemoteAddr().String())
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			return
		}
		str := string(buf[:n])
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "收到客户端的消息",
			Content: "Content:" + str,
		})
		data.TcpSSet.AcceptMsgList = append(data.TcpSSet.AcceptMsgList, str+"\n")
	}
}

func TcpServerSendData(data string) {
	data = strings.TrimSpace(data)
	_, err := TcpServerConn.Write([]byte(data))
	if err != nil {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Server发送失败",
			Content: "Content:" + err.Error(),
		})
	} else {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Server发送成功",
			Content: "",
		})
	}
}
func TcpServerClose(Ch chan net.Conn) {
	tcpServerConn := <-Ch
	err := tcpServerConn.Close()
	if err != nil {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "server关闭失败",
			Content: "Content:" + err.Error(),
		})
	} else {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "server关闭成功",
			Content: "",
		})
	}
}

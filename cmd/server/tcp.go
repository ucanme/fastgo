package server

import (
	"context"
	"fmt"
	"github.com/smallnest/goframe"
	logger "github.com/ucanme/fastgo/library/log"
	"net"
	"time"
)

var CmdChan  = make(chan string)

func StartTcpClient()  {
	var data []byte
	var err error
	var conn net.Conn
	for {
		ctx,cancel := context.WithCancel(context.Background())
		conn,err = net.Dial("tcp","127.0.0.1:8082")
		if err != nil{
			logger.LogError(map[string]interface{}{"tcp client dail fail" : time.Now().Unix()})
			time.Sleep(3*time.Second)
		}
		c := goframe.NewLineBasedFrameConn(conn)
		//消息收取
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					data, err = c.ReadFrame()
					if err != nil {
						cancel()
						return
					}
					fmt.Println("-----------------------", string(data))
				}
			}
		}(ctx)


		go func(ctx context.Context) {
			select {
				case <- ctx.Done():
					return
			case msg := <- CmdChan:
				fmt.Println(msg)
				err = c.WriteFrame([]byte(msg))
				if err != nil{
					cancel()
					return
				}
			}
		}(ctx)
}}

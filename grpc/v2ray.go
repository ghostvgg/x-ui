package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"v2ray.com/core/app/proxyman/command"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/vmess"
)

const (
	API_ADDRESS = "46.101.211.210"
	API_PORT    = 10085
	INBOUND_TAG = "proxy"
	LEVEL       = 0
	EMAIL       = "yuser"
	UUID        = "3f278b0e-c8c5-11ed-afa1-0242ac120002"
	ALTERID     = 64
)

func addUser(c command.HandlerServiceClient) {
	resp, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: INBOUND_TAG,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Level: LEVEL,
				Email: EMAIL,
				Account: serial.ToTypedMessage(&vmess.Account{
					Id:               UUID,
					AlterId:          ALTERID,
					SecuritySettings: &protocol.SecurityConfig{Type: protocol.SecurityType_AUTO},
				}),
			},
		}),
	})
	if err != nil {
		log.Printf("failed to call grpc command: %v", err)
	} else {
		log.Printf("ok: %v", resp)
	}
}
func removeUser(c command.HandlerServiceClient) {
	resp, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: INBOUND_TAG,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: EMAIL,
		}),
	})
	if err != nil {
		log.Printf("failed to call grpc command: %v", err)
	} else {
		log.Printf("ok: %v", resp)
	}
}
func main() {
	cmdConn, err := grpc.Dial(fmt.Sprintf("%s:%d", API_ADDRESS, API_PORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	hsClient := command.NewHandlerServiceClient(cmdConn)
	addUser(hsClient)
	// removeUser(hsClient)
}

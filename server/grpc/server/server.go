package server

import (
	"IM/protobuf"
	"IM/server/websocket"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) SendMsgAll(c context.Context, req *protobuf.SendMsgAllReq) (rsp *protobuf.SendMsgAllRsp, err error) {
	fmt.Println("grpc_request 给本机全体用户发消息", req.String())

	rsp = &protobuf.SendMsgAllRsp{}

	data := protobuf.GetMsgData("msg", req.GetUid(), req.GetSeq(), req.GetMessage())
	pushId, err := websocket.SendMessageToUsers(req.GetUid(), req.AppId, req.GetTargetIds(), data)

	rsp.PushIds = pushId

	fmt.Println("grpc_response 给本机全体用户发消息:", rsp.String())
	return
}

func Init() {
	rpcPort := viper.GetString("app.rpcPort")
	fmt.Println("rpc server 启动", rpcPort)

	lis, err := net.Listen("tcp", ":"+rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterIMServerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

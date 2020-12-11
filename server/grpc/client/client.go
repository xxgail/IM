package client

import (
	"IM/models"
	"IM/protobuf"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func SendMsgAll(server *models.Server, appId string, uid string, targetIds []string, seq string, message string) (pushId []string, err error) {
	log.SetOutput(os.Stdout)

	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		log.Fatalln()
	}
	defer conn.Close()

	c := protobuf.NewIMServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &protobuf.SendMsgAllReq{
		AppId:     appId,
		Uid:       uid,
		TargetIds: targetIds,
		Seq:       seq,
		Message:   message,
	}

	rsp, err := c.SendMsgAll(ctx, req)
	if err != nil {
		log.Fatalf("登陆失败！%v", err)
		return
	}

	pushId = rsp.PushIds

	fmt.Println("给全体用户发送消息 成功!")

	return
}

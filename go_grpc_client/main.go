package main

import (
	"context"
	"go_grpc_client/app"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:5000"
	defaultName = "notify"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	client := app.NewPushNotificationClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SendNotification(context.Background(), &app.NotificationData{
		Message: "this is a message",
		Code:    "FR234",
		Payload: "{ 'id': '1', 'name': 'pawn'}",
	})
	if err != nil {
		log.Panic(err)
	}
	log.Print("send notification success: ", res.Success)
}

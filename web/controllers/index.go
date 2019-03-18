package controllers

import (
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"fmt"
	"github.com/ghoiufyia/WxApp.kindle/web/models"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "log"
)

type IndexController struct {
	dogo.Controller
	Name	string
	age		int32
}

func (i *IndexController)Json ()  {
	db := dogo.GetDB()
	var user_email models.UserEmail
	db.First(&user_email)

	i.SetData("code","1")
	i.SetData("msg","ok")
	// i.SetData("data",make(map[string]string, 0))
	i.SetData("data",user_email)
	i.RenderJson()
}


func (i *IndexController)Index ()  {
	fmt.Printf("adsd===============================")
	// fmt.Println(i.Ctx)
	// fmt.Printf("%+v",i.Ctx.Request)

	fmt.Printf("adsd========ffffff=======================")

	i.Render()

	// // 连接远端服务
	// conn,err := grpc.Dial(ADDRESS,grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("connect error %v",err)
	// }
	// defer conn.Close()
	// // 定义客户端
	// client := email_pb.NewEmailServiceClient(conn)

	// var resuest = email_pb.CreateEmailRequest{UserId:"fghfyffffffffjy",Email:"ddd@11.com"}

	// // 调用 RPC
	// resp, err := client.CreateEmail(context.Background(), &resuest)
	// if err != nil {
	// 	log.Printf("create email error: %v", err)
	// }

	// log.Printf("created: %t", resp.Msg)
	// conn.Close()

}


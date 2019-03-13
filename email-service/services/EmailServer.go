package services

import (
	email_pb "github.com/ghoiufyia/WxApp.kindle/email-service/proto/email"
	"context"
	"github.com/ghoiufyia/WxApp.kindle/email-service/models"
	"github.com/jinzhu/gorm"
	"time"
	"github.com/RichardKnop/uuid"
	"errors"
	"fmt"
)


//邮件服务数据结构，实现EmailServiceServer
type EmailServer struct {
	Db *gorm.DB
}

func (s *EmailServer)CreateEmail(ctx context.Context, req *email_pb.CreateEmailRequest) (*email_pb.CreateEmailResponse, error){
	user_id := req.GetUserId()
	if user_id == "" {
		return nil,errors.New("UserId Not Found.")
	}
	email := req.GetEmail()
	if email == "" {
		return nil,errors.New("Email Not Found.")
	}
	user_email := models.UserEmail{
		BaseModel:models.BaseModel{
			ID:			uuid.New(),
			CreatedAt: 	time.Now().Unix(),
			UpdatedAt: 	time.Now().Unix(),
			DeletedAt: 	0,
		},
		UserId:user_id,
		Email:email,
	}
	fmt.Println("%v",user_email)
	if err := s.Db.Create(user_email).Error; err != nil {
		return nil, err
	}
	rep := &email_pb.CreateEmailResponse{
		Code:	1,
		Msg:	"添加成功",
	}
	return rep,nil
}

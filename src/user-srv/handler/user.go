package handler

import (
	"MIX_GRPC/src/share/errors"
	"MIX_GRPC/src/user-srv/db"
	"context"
	"go.uber.org/zap"
	"MIX_GRPC/src/share/pb"
)

type UserServiceHandler struct {
	logger *zap.Logger
}

// 注册用户
func (u *UserServiceHandler) RegistAccount(ctx context.Context, req *pb.RegistAccountReq, rsp *pb.RegistAccountRsp) error {
	uname := req.Uname
	password := req.Password
	email := req.Email
	user, err := db.SelectUserByEmail(email)
	if err != nil && user == nil{
		return errors.ErrorUserFailed
	}else if err == nil && user != nil{
		return errors.ErrorUserAlready
	}
	err = db.InsertUser(uname, password, email)
	if err != nil{
		return errors.ErrorUserFailed
	}
	return nil
}

//通过id查询某用户信息
func (u *UserServiceHandler) GetUinfoByEmail(ctx context.Context, req *pb.GetUinfoByEmailReq, rsp *pb.GetUinfoByEmailRsp) error {
	email := req.Email
	user, err := db.SelectUserByEmail(email)
	if err != nil && user == nil{
		return errors.ErrorUserFailed
	} else if user == nil{
		return errors.ErrorUserNotExists
	}
	rsp.Id = user.Id
	rsp.Uname = user.UName
	rsp.Email = user.Email
	rsp.Crtime = user.CrTime
	return nil
}

//获取所有用户信息
func (u *UserServiceHandler) GetAllUinfo(ctx context.Context, req *pb.GetAllUinfoReq, rsp *pb.GetAllUinfoRsp) error{
	users, err := db.SelectAllUser()
	if err != nil{
		return errors.ErrorUserFailed
	}
	for _, row := range users{
		user := pb.GetUinfoByEmailRsp{}
		user.Id = row.Id
		user.Uname = row.UName
		user.Email = row.Email
		user.Crtime = row.CrTime
		rsp.Alluinfo = append(rsp.Alluinfo, &user)
	}
	return nil
}
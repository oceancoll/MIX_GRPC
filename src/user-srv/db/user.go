package db

import (
	"MIX_GRPC/src/user-srv/entity"
	"database/sql"
	"fmt"
	"time"
)

func SelectUserByEmail(email string) (*entity.User, error) {
	user := entity.User{}
	err := db.Get(&user, "select * from user where email=?", email)
	if err == sql.ErrNoRows{
		return nil, nil
	} else if err != nil{
		return nil, err
	}
	return &user, nil
}

func InsertUser(uname string, password string, email string) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("insert into `user` (`uname`, `password`, `crtime`, `email`) values (?,?,?,?)", uname, password, now, email)
	return err
}

func SelectAllUser() ([]entity.User, error)  {
	var users []entity.User
	err := db.Select(&users, "select * from user order by `id`")
	fmt.Println(err)
	if err != nil{
		return nil, err
	}
	return users, nil
}
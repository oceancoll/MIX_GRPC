package entity

//用户定义数据表的格式与代码层的映射关系

type User struct {
	Id int32 `json:"id" db:"id"`
	UName string `json:"uname" db:"uname"`
	Password string `json:"password" db:"password"`
	CrTime string `json:"crtime" db:"crtime"`
	Email string `json:"email" db:"email"`
	Gender int8 `json:"gender" db:"gender"`
	Phone string `json:"phone" db:"phone"`
}

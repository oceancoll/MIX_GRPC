package entity

//购买信息定义数据表的格式与代码层的映射关系

type Sale struct {
	Id int32 `json:"id" db:"id"`
	UId int32 `json:"uid" db:"uid"`
	Itemname string `json:"itemname" db:"itemname"`
	CrTime string `json:"crtime" db:"crtime"`
	Price float32 `json:"price" db:"price"`
}

type SaleItem struct {
	Uname string `json:"uname"`
	Itemname string `json:"itemname"`
	Price float32 `json:"price"`
	CrTime string `json:"crtime"`
}

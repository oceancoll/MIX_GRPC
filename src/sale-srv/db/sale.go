package db

import (
	"MIX_GRPC/src/sale-srv/entity"
	"time"
)

func InsertSaleItem(uid int32, itemname string, price float32) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("insert into `sale` (`uid`, `itemname`, `crtime`, `price`) values (?,?,?,?)", uid, itemname, now, price)
	return err
}

func GetBuyitemsByUid(uid int32) ([]entity.Sale, error) {
	var sales []entity.Sale
	err := db.Select(&sales, "select * from sale where uid=? order by `crtime` desc ", uid)
	if err != nil{
		return nil, err
	}
	return sales, nil
}

func SelectAllItems() ([]entity.SaleItem, error) {
	var items []entity.SaleItem
	err := db.Select(&items, "select a.itemname, a.crtime, a.price, b.uname from sale a inner join user b on a.uid = b.id")
	if err != nil{
		return nil, err
	}
	return items, nil
}

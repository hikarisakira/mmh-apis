package models

import "time"

// oracledb@odrh_del(table@mmh)
type SearchFormat struct {
	Pno int `json:"pno" binding:"required" xorm:"notnull number(8)"`
}
type FeedBackFormat struct {
	Pno     int       `json:"pno" binding:"required" xorm:"number(8) notnull"`
	SchDate time.Time `json:"schdate" binding:"required xorm:"date notnull"`
	Dr      string    `json:"dr" binding:"required" xorm:"varchar2(4) notnull"`
	OperDhm time.Time `json:"oper_dhm" binding:"required" xorm:"date notnull"`
	Memo    string    `json:"memo" binding:"required" xorm:"varchar2(100) notnull"`
	Udate   time.Time `json:"udate" binding:"required" xorm:"date notnull"`
	Xuser   string    `json:"xuser" binding:"required" xorm:"varchar2(4) notnull"`
}

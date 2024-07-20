package models

// oracledb@odrh_del(table@mmh)
type FeedBackFormat struct {
	An_Code  string `json:"an_Code" binding:"required" xorm:"char(7) notnull"`
	An_Name  string `json:"an_Name" binding:"required" xorm:"char(20) notnull"`
	An_Cout1 string `json:"an_Cout1"  xorm:"char(20) notnull"`
	An_Cout2 string `json:"an_Cout2"  xorm:"char(20) notnull"`
	An_Cout3 string `json:"an_Cout3"  xorm:"char(20) notnull"`
	An_Cout4 string `json:"an_Cout4"  xorm:"char(20) notnull"`
	An_Cout5 string `json:"an_Cout5"  xorm:"char(20) notnull"`
	An_Cout6 string `json:"an_Cout6"  xorm:"varchar2(100) notnull"`
}

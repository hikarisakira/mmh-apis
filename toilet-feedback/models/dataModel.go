package models

// oracledb@odrh_del(table@mmh)
type FeedBackFormat struct {
	PlaceCode string `json:"placeCode" binding:"required" xorm:"char(7) notnull"`
	PlaceName string `json:"placeName" binding:"required" xorm:"char(50) notnull"`
	An_Cout1  string `json:"an_Cout1"  xorm:"char(20) notnull"`
	An_Cout2  string `json:"an_Cout2"  xorm:"char(20) notnull"`
	An_Cout3  string `json:"an_Cout3"  xorm:"char(30) notnull"`
	An_Cout4  string `json:"an_Cout4"  xorm:"varchar2(100) notnull"`
	Bn_Cout1  string `json:"bn_Cout1"  xorm:"char(20) notnull"`
	Bn_Cout2  string `json:"bn_Cout2"  xorm:"char(20) notnull"`
	Bn_Cout3  string `json:"bn_Cout3"  xorm:"char(20) notnull"`
	Bn_Cout4  string `json:"bn_Cout4"  xorm:"char(20) notnull"`
	Bn_Cout5  string `json:"bn_Cout5"  xorm:"char(20) notnull"`
	Bn_Cout6  string `json:"bn_Cout6"  xorm:"char(20) notnull"`
	Bn_Cout7  string `json:"bn_Cout7"  xorm:"varchar2(100) notnull"`
}

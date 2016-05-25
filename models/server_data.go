// server_data
package models

type ShopCarData struct {
	Goodsid    int64
	Name       string
	Count      int32
	Price      int32
	Createdate string
	Imgsrc     string
	Tourl      string
}

type ShopGoodsinfo struct {
	Goodsid int64
	Price   int32
	Name    string
	Imgsrc  string
	Tourl   string
}

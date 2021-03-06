// sbtatruct.go
package models

type Character struct {
	Uid           int64 `orm:"pk;auto"`
	Uname         string
	Nickname      string `orm:"size(20);default(default)"`
	Pw            string
	Icon          string `orm:"null"`
	Phone         string `orm:null;size(11)`
	Point         int32  `orm:"default(0)"`
	Money         int32  `orm:"default(0)"`
	Qqnum         int64  `orm:"default(0)"`
	Type          int8
	Email         string `orm:"null"`
	IdCard        string `orm:null;size(18)`
	Local         string `orm:"null"`
	CreateTime    string
	LastloginTime string `orm:"null"`
	Token         string
	Platform      string `orm:"null"`
}

type Goodsinfo struct {
	Goodsid         int64  `orm:"pk;auto"`
	Name            string `orm:"size(50)"`
	Price           int32
	Discount        int32  `orm:"default(0)"`
	Barcode         int64  `orm:"default(0)"`
	Updatetime      string `orm:"null"`
	Createtime      string
	Editer          string
	Ishot           int32  `orm:"default(0)"`
	Ishave          int32  `orm:"default(0)"`
	Count           int32  `orm:"default(0)"`
	Content         string `orm:"null;size(200)"`
	DiscountEndtime string `orm:"null"`
	Categoryid      int32  `orm:"null"`
	CategoryName    string `orm:null;size(20)`
	Couponid        int64  `orm:"null"`
	Tourl           string `orm:"null"`
}

type Orderbook struct {
	Id          int64 `orm:"pk;auto"`
	Bid         int64
	Uid         int64
	Goodsid     string `orm:size(500)`
	Addressid   int64
	Paytype     int8
	Costpoint   int64
	Postage     int32
	Titlemoney  int64
	Paymoney    int64
	Couponmoney int32
	Bookingdate string
	PostDate    string
	Status      int8
}

type Picture struct {
	Imgid      int64  `orm:"pk;auto"`
	Goodsid    int64  `orm:"null"`
	Imgsrc     string `orm:"size(1000)"`
	Type       string `orm:"size(20);null"` //图片作用
	Where      string `orm:"size(20);null"` //什么模块使用
	Isshow     int8   `orm:"default(0)"`
	Updatetime string `orm:"null"`
	Createtime string
	Tourl      string `orm:"null"`
}

type EverydaySales struct {
	Id      int64 `orm:"pk;auto"`
	Goodsid int64
	Value   int32 //每日销量
	Date    string
}

type Evaluate struct {
	Id         int64 `orm:"pk;auto"`
	Goodsid    int64
	Uname      string `orm:"null;size(20)"`
	Uid        int64
	PictureSrc string `orm:"null;size(1000)"`
	Createtime string
	Good_star  int8 `orm:"default(0)"`
	Post_star  int8 `orm:"default(0)"`
}

type Rechage struct {
	Id      int64 `orm:"pk;auto"`
	Orderid int64
	Uid     int64
	Type    int8
	Value   int32
}

type Collect struct {
	Id      int64 `orm:"pk;auto"`
	Uid     int64
	Goodsid int64
}

type DiscountCoupon struct {
	Id         int64 `orm:"pk;auto"`
	Couponid   int64
	Uid        int64
	Count      int32
	Price      int32
	Starttime  string
	Endtime    string
	CouponType int32
}

type ExpenseLog struct {
	Id    int64 `orm:"pk;auto"`
	Uid   int64
	Type  int8
	Money int64
	Date  string
	Bkid  int64
	Point int64
}

type Online struct {
	Id      int64 `orm:"pk;auto"`
	Uid     int64
	Uname   string
	Type    int8
	Statues int8
	Token   string
}

type Address struct {
	Id       int64 `orm:"pk;auto"`
	Uid      int64
	Name     string
	Location string `orm:"null;size(100)"`
	Fulladdr string `orm:"null;size(200)"`
	PhonNum  string `orm:"null;size(20)"`
	Isdef    int8   `orm:default(0)`
}

type Poster struct {
	Posterid int32  `orm:"pk;auto"`
	Name     string `orm:"size(20)"`
	Phone    string `orm:size(11)`
	Postgift int64
}

type Postergift struct {
	Id       int64 `orm:"pk;auto"`
	Posterid int32
	Giftid   int32
	Giftname string
	Count    int32
	Point    int64
}

type GiftLog struct {
	Id       int64
	Uid      int64
	Posterid int32
	Giftid   int32
	Giftname string
	Count    int32
	Point    int64
}

type Shopcar struct {
	Id         int64
	Uid        int64
	Goodsid    int64
	Count      int32
	CreateDate string
}

type Category struct {
	Id         int32 `orm:"pk;auto"`
	Categoryid int32
	Name       string
}

// sbtatruct.go
package models

type Character struct {
	Uid           int64  `orm:"pk;auto"`
	Uname         string `orm:"size(20)"`
	Pw            string
	Icon          string `orm:"null"`
	Phone         string `orm:null;size(11)`
	Point         int32  `orm:"default(0)"`
	Money         int32  `orm:"default(0)"`
	Email         string `orm:"null"`
	IdCard        string `orm:null;size(18)`
	Local         string `orm:"null"`
	CreateTime    string
	LastloginTime string `orm:"null"`
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
	Categorytype    int32  `orm:"null"`
	CategoryName    string `orm:null;size(20)`
	Couponid        int64  `orm:"null"`
	Tourl           string `orm:"null"`
}

type Book struct {
	Bkid       int64 `orm:"pk;auto"`
	Uid        int64
	Goodsid    int64
	Name       string `orm:"size(50)"`
	Count      int32
	Addr       string `orm:"size(100)"`
	Statues    int16
	Price      int32
	Postage    int32 `orm:"default(0)"`
	Createtime string
	Endtime    string `orm:"null"`
}

type Picture struct {
	Imgid      int64  `orm:"pk;auto"`
	Goodsid    int64  `orm:"null"`
	Imgsrc     string `orm:"size(1000)"`
	Type       string `orm:"size(20);null"` //图片作用
	Where      string `orm:"size(20);null"`
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
	Id         int64 `orm:"pk;auto"`
	Uid        int64
	Uname      string
	Type       int8
	Statues    int8
	SessionKey string
}

type Address struct {
	Id       int64 `orm:"pk;auto"`
	Uid      int64
	Name     string
	Addr     string `orm:"null;size(1000)"`
	Postcode string `orm:"null;size(6)"`
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

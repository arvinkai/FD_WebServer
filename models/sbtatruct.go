// sbtatruct.go
package models

type Character struct {
	Uid           int64  `orm:"pk;auto"`
	Uname         string `orm:"size(20)"`
	Pw            string
	Icon          string
	Phone         string `orm:size(11)`
	Point         int32
	Money         int32
	Email         string
	IdCard        string `orm:size(18)`
	Local         string
	CreateTime    uint64
	LastloginTime uint64
}

type Goodsinfo struct {
	Goodsid         int64  `orm:"pk;auto"`
	Name            string `orm:"size(50)"`
	Price           int32
	Discount        int32
	Barcode         int64
	Updatetime      string
	Createtime      string
	Editer          string
	Ishot           int32
	Ishave          int32
	Count           int32
	Content         string `orm:"size(200)"`
	DiscountEndtime string
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
	Postage    int32
	Createtime string
	Endtime    string
}

type Picture struct {
	Imgid      int64 `orm:"pk;auto"`
	Goodsid    int64
	Imgsrc     string `orm:"size(1000)"`
	Type       int16  //图片作用
	Isshow     int16
	Updatetime string
	Createtime string
}

type Category struct {
	Goodsid      int64 `orm:"pk"`
	Categoryid   int32 `orm:"pk"`
	Categorytype int32
	CategoryName string `orm:size(20)`
	Couponid     int64
}

type EverydaySales struct {
	Goodsid int64
	Value   int32 //每日销量
	Date    string
}

type Evaluate struct {
	Goodsid    int64
	Uname      string `orm:"size(20)"`
	Uid        int64
	PictureSrc string `orm:"size(1000)"`
	Createtime string
	Good_star  int8
	Post_star  int8
}

type Poster struct {
	Postid   int32  `orm:"pk;auto"`
	Name     string `orm:"size(20)"`
	Phone    string `orm:size(11)`
	Postgift int64
}

type Rechage struct {
	Id      int64 `orm:"pk;auto"`
	Orderid int64 `orm:"pk"`
	Uid     int64
	Type    int8
	Value   int32
}

type Collect struct {
	Uid     int64 `orm:"pk"`
	Goodsid int64 `orm:"pk"`
}

type DiscountCoupon struct {
	Couponid   int64
	Uid        int64
	Count      int32
	Starttime  string
	Endtime    string
	CouponType int32
}

type ExpenseLog struct {
	Uid   int64
	Type  int8
	Money int64
	Date  string
	Bkid  int64
	Point int64
}

type Online struct {
	Uid        int64
	Uname      string
	Type       int8
	Statues    int8
	SessionKey string
}

type Address struct {
	Uid      int64
	Name     string
	Addr     string `orm:"size(1000)"`
	Postcode string `orm:"size(6)"`
}

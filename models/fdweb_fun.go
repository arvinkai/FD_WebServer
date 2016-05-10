// fdweb_fun
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func GetPicture(Type string) (resp []*Picture) {
	fmt.Println(Type)
	var ptc []*Picture = make([]*Picture, 0)
	switch Type {
	case "Home":
		{
			o := orm.NewOrm()
			qs := o.QueryTable("picture")
			_, err := qs.Filter("type", 1).Filter("isshow", 1).All(&ptc)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	case "shop":
		{

		}
	}
	return ptc
}

func GetGoodsInfo(Type string) string {
	return "test"
}

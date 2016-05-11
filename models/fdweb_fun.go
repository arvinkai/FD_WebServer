// fdweb_fun
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func GetPicture(Where string, Type string, IsShow int8) (resp []*Picture, cols int64, err error) {
	fmt.Println(Type)
	var ptc []*Picture = make([]*Picture, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("picture")
	a, err := qs.Filter("type", Type).Filter("isshow", IsShow).Filter("where", Where).All(&ptc)
	if err != nil {
		fmt.Println(err)
		return nil, a, err
	}
	if a == 0 {
		return nil, a, nil
	}

	return ptc, a, nil
}

func GetGoodsInfo(Type string) string {
	return "test"
}

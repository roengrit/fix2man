package models

import "github.com/astaxie/beego/orm"

//Topic _
type Topic struct {
	ID    int
	Name  string `orm:"size(225)"`
	Mark1 int
	Mark2 int
	Mark3 int
}

//TopicAssess _
type TopicAssess struct {
	ID     int
	DocNo  string
	Remark string `orm:"size(225)"`
}

//TopicAssessSub _
type TopicAssessSub struct {
	ID    int
	DocNo string
	Name  string `orm:"size(225)"`
	Mark1 int
	Mark2 int
	Mark3 int
}

func init() {
	orm.RegisterModel(
		new(Topic),
		new(TopicAssess),
		new(TopicAssessSub),
	) // Need to register model in init
}

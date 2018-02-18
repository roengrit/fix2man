package models

import (
	"github.com/astaxie/beego/orm"
)

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
	Image  string
}

//TopicAssessSub _
type TopicAssessSub struct {
	ID      int
	TopicID int
	DocNo   string
	Name    string
	Mark1   int
	Mark2   int
	Mark3   int
	Mark    int
	Flag    int
}

func init() {
	orm.RegisterModel(
		new(Topic),
		new(TopicAssess),
		new(TopicAssessSub),
	) // Need to register model in init
}

//GetAllTopic _
func GetAllTopic(DocNo string) (req *[]TopicAssessSub) {
	o := orm.NewOrm()
	reqGet := &[]TopicAssessSub{}
	o.Raw(`
		SELECT
		topic.i_d,
		topic.name,
		topic.mark1,
		topic.mark2,
		topic.mark3,
		topic_assess_sub.doc_no,
		topic_assess_sub.mark,
		topic_assess_sub.flag
	FROM
		topic
		LEFT JOIN topic_assess_sub ON topic.i_d = topic_assess_sub.topic_i_d and topic_assess_sub.doc_no = '` + DocNo + `'`).QueryRows(reqGet)
	return reqGet
}

package models

import (
	"../mongo"
	"fmt"
	"github.com/melonws/goweb/libs/logHelper"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type PhotoItem struct {
	Title    string `bson:"title"`
	Imgs     string `bson:"imgs"`
	CreateAt int64  `bson:"create_at"`
}
type return_photo_item struct {
	title    string `json:"title"`
	imgs     string `json:"imgs"`
	createAt int64  `json:"create_at"`
}

//添加一条记录
func AddItem(t *PhotoItem) (status int64, err error) {
	println(t.Title)
	session, connection := mongo.CreateModel("photo_wall_list")
	if session == nil {
		return 0, errors.New("没连上啊")
	}
	defer session.Close()
	data := &PhotoItem{
		Title:    t.Title,
		Imgs:     t.Imgs,
		CreateAt: t.CreateAt,
	}
	fmt.Println(data)
	error := connection.Insert(data)
	if error != nil {
		logHelper.WriteLog("[插入数据失败]", "mongo/error")
		return 0, error
	}
	logHelper.WriteLog("[插入数据成功]", "mongo/notify")
	return 1, nil
}

func GetList() (status int, data interface{}, err error) {
	session, connection := mongo.CreateModel("photo_wall_list")
	if session == nil {
		return 0, nil, errors.New("没连上啊")
	}
	defer session.Close()
	query := []bson.M{}
	// 接受数据结果（不能用Result直接定义结果，感觉可以，但是不能实现，有待理解
	var values []interface{}
	//筛选字段,_id主键默认返回，如不需要返回id，需要显式定义_id为0
	filter := bson.M{"title": 1, "imgs": 1, "createAt": 1}
	error := connection.Find(query).Select(filter).All(&values)
	println("返回值")
	fmt.Println(values)

	if error != nil {
		logHelper.WriteLog("获取数据失败"+error.Error(), "mongo/error")
		return 0, nil, errors.New("获取数据失败")
	}
	return 1, values, nil
}

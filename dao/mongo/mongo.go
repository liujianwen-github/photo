package mongo

import (
"github.com/melonws/goweb/libs/logHelper"
"gopkg.in/mgo.v2"
"time"
)

func connect() *mgo.Session {

	mongoDialInfo := &mgo.DialInfo{
		Addrs:[]string{"47.104.7.232:27017"},
		Direct:false,
		Timeout:time.Second*5,
		Database:"photowall",
		Username:"admin",
		Password:"123456",
	}
	//mongoDialInfo,_ := mgo.ParseURL("mongodb://root:root@47.104.7.232:27017/memo")
	//tlsConfig := &tls.Config{}
	//mongoDialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
	//	conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
	//	return conn, err
	//}
	session, err := mgo.DialWithInfo(mongoDialInfo)

	if err != nil {
		logHelper.WriteLog("mongo 连接失败"+err.Error(), "mongodb/error")
		return nil
	}
	logHelper.WriteLog("mongo 连接成功", "mongo/access")
	return session
}

/**
 * 通过数据库连接后，拿到对应的session(会话) 和 collection(集合)
 * 便于操作
 * 返回后的操作，记得先 defer session.Close()
 */
func CreateModel(collectionName string) (*mgo.Session, *mgo.Collection) {

	session := connect()

	if session == nil {
		return nil, nil
	}

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("photowall").C(collectionName)

	return session, collection
}

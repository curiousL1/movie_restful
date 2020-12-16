package models

import (
	"log"
	//"fmt"
	"github.com/globalsign/mgo"
)

const (
	HOST   = "127.0.0.1:27017"
	DBSOURCE = "Movies"
	USERNAME   = "service_test"
	PASSWORD   = "123456"
)

var globalS *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{HOST},
		Source:   DBSOURCE,
		Username: USERNAME,
		Password: PASSWORD,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error ", err)
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	mgoSess := globalS.Copy()
	mgoCollec := mgoSess.DB(db).C(collection)
	return mgoSess, mgoCollec
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, mc := connect(db, collection)
	defer ms.Close()
	return mc.Insert(docs...)
}

func Find(db, collection string, query, selector, result interface{}) error {
	ms, mc := connect(db, collection)
	defer ms.Close()
	return mc.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, mc := connect(db, collection)
	defer ms.Close()
	return mc.Find(query).Select(selector).All(result)
}

func Update(db, collection string, query, update interface{}) error {
	ms, mc := connect(db, collection)
	defer ms.Close()
	return mc.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, mc := connect(db, collection)
	defer ms.Close()
	return mc.Remove(query)
}

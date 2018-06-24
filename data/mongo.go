package data

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

type MongoDBConn struct {
	session *mgo.Session
	db      string
}

func NewMongoDBConn() *MongoDBConn {
	return &MongoDBConn{}
}

/*func (m *MongoDBConn) Connect(hosts []string, userName, password, replicaSet,
	dbName, authDb string, mode mgo.Mode, refresh bool) *mgo.Session {

	hostsString := strings.Join(hosts, ",")
	var url string
	if len(userName) == 0 && len(password) == 0 && len(replicaSet) == 0 {
		url = "mongodb://" + hostsString
	} else {
		url = "mongodb://" + userName + ":" + password + "@" + hostsString + "/" + authDb + "?replicaSet=" + replicaSet
	}

	dialinfo, err := mgo.ParseURL(url)
	if err != nil {
		log.Println("Couldn't parse mongodb url ", url)
		log.Fatalln(err)
	}

	session, err := mgo.Dial(url)
	if err != nil {
		log.Println("URL: ", url)
		log.Printf("Couldn't connect to %v", dialinfo.Database)
		log.Fatalln(err)
	}

	m.SetDB(dbName)

	session.SetMode(mode, refresh)
	session.SetSocketTimeout(time.Duration(10 * time.Minute))

	m.session = session
	return m.session
}*/

func (m *MongoDBConn) Connect(host string, dbName string) *mgo.Session {

	var url string

	url = "mongodb://" + host

	dialinfo, err := mgo.ParseURL(url)
	if err != nil {
		log.Println("Couldn't parse mongodb url ", url)
		log.Fatalln(err)
	}

	session, err := mgo.Dial(url)
	if err != nil {
		log.Println("URL: ", url)
		log.Printf("Couldn't connect to %v", dialinfo.Database)
		log.Fatalln(err)
	}

	m.SetDB(dbName)

	//session.SetMode(mode, refresh)
	session.SetSocketTimeout(time.Duration(10 * time.Minute))

	m.session = session
	return m.session
}

func (m *MongoDBConn) SetDB(db string) {
	m.db = db
}

func (m *MongoDBConn) GetDB() string {
	return m.db
}

func (m *MongoDBConn) Stop() {
	m.session.Close()
}

func (m *MongoDBConn) GetConn() *mgo.Session {
	return m.session.Copy()
}

func (m *MongoDBConn) SetIndex(key, db, collection string) error {
	index := mgo.Index{
		Key:        []string{key},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	collectionDb := m.session.DB(db).C(collection)
	err := collectionDb.EnsureIndex(index)
	return err
}

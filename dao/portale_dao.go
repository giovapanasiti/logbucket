package dao

import (
	"log"

	"gopkg.in/mgo.v2"
	"github.com/giovapanasiti/logbucket/models"
)

type LogPortaleDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "logportale"
)

// Establish a connection to database
func (m *LogPortaleDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//// Find list of logs
//func (m *LogPortaleDAO) FindAll() ([]models.LogPortale, error) {
//	var logsportale []models.LogPortale
//	err := db.C(COLLECTION).Find(bson.M{}).All(&logsportale)
//	return logsportale, err
//}
//
//// Find a movie by its id
//func (m *LogPortaleDAO) FindById(id string) (models.LogPortale, error) {
//	var logportale models.LogPortale
//	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&logportale)
//	return logportale, err
//}

// Insert a logportale into database
func (m *LogPortaleDAO) Insert(logportale models.LogPortale) error {
	err := db.C(COLLECTION).Insert(&logportale)
	return err
}

//// Delete an existing logportale
//func (m *LogPortaleDAO) Delete(logportale models.LogPortale) error {
//	err := db.C(COLLECTION).Remove(&logportale)
//	return err
//}
//
//// Update an existing logportale
//func (m *LogPortaleDAO) Update(logportale models.LogPortale) error {
//	err := db.C(COLLECTION).UpdateId(logportale.ID, &logportale)
//	return err
//}
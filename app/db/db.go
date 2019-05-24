package db

import (
	"github.com/globalsign/mgo"
)

//Sess .
var Sess *mgo.Session

//DB .
var DB *DBConfig

//DBConfig .
type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"databases"`
	Auth     Auth
}

//Auth .
type Auth struct {
	Mechanism string `json:"mechanism"` //.
	Db        string `json:"db"`        //auth db.
}

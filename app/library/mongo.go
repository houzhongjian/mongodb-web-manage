package library

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"dbclient.net/web/mongodb/app/db"
)

//Mgo .
type Mgo struct {
}

//Cmd .
func (m Mgo) cmd(cmd string) (string, error) {
	cf := db.DB
	sh := fmt.Sprintf("mongo %s/%s --authenticationDatabase '%s' --username '%s' --password '%s' --quiet --authenticationMechanism '%s'  --eval=\"%s\";",
		cf.Host,
		cf.Database,
		cf.Auth.Db,
		cf.User,
		cf.Password,
		cf.Auth.Mechanism,
		cmd)
	log.Println("sh:", sh)
	b, err := exec.Command("bash", "-c", sh).Output()
	if err != nil {
		log.Printf("%+v\n", err)
		return "", err
	}
	js := string(b)
	log.Println("source shell:", js)

	if strings.Contains(js, "WriteResult") {
		return js, nil
	}

	if strings.Contains(js, "{") {
		js = "[" + js + "]"
		js = strings.Replace(js, "\n", "", -1)
		js = strings.Replace(js, "ISODate(", "", -1)
		js = strings.Replace(js, "ObjectId(", "", -1)
		js = strings.Replace(js, "NumberLong(", "", -1)
		js = strings.Replace(js, "UUID(", "", -1)
		js = strings.Replace(js, "Type \"it\" for more", "", -1)
		js = strings.Replace(js, "}{", "},{", -1)
		js = strings.Replace(js, "\")", "\"", -1)
		js = strings.Replace(js, ")", "", -1)
	}
	return js, nil
}

//DbList .
// func (m Mgo) DbList() []db.Database {
// 	js, err := m.cmd("db.adminCommand('listDatabases')")
// 	if err != nil {
// 		log.Printf("%+v\n", err)
// 		return nil
// 	}
// 	dbList := db.Databases{}
// 	if err := json.Unmarshal([]byte(js), &dbList); err != nil {
// 		log.Printf("%+v\n", err)
// 		return nil
// 	}
// 	return dbList.Database
// }

//CollList .
// func (m Mgo) CollList() []string {
// 	js, err := m.cmd("db.getCollectionNames()")
// 	if err != nil {
// 		log.Printf("%+v\n", err)
// 		return nil
// 	}
// 	// js = strings.Replace(js, `"id" : NumberLong(0),`, "", -1)
// 	// js = strings.Replace(js, `"uuid" : UUID("5c19dadd-3107-4352-807d-c502a133f863"),`, "", -1)
// 	log.Println(js)
// 	collList := []string{}
// 	if err := json.Unmarshal([]byte(js), &collList); err != nil {
// 		log.Printf("%+v\n", err)
// 		return nil
// 	}
// 	return collList

// }

//Query .
func (m Mgo) Query(cmd string) string {
	js, err := m.cmd(cmd)
	if err != nil {
		log.Printf("%+v\n", err)
		return ""
	}

	return js
}

//ParseShell .
func (m Mgo) ParseShell(shell string) string {
	if m.isShowDbs(shell) {
		return `db.adminCommand('listDatabases')`
	}

	if m.isShowCollections(shell) {
		return `db.adminCommand('listCollections')`
	}

	if m.isAggregate(shell) {
		shell = strings.Replace(shell, ` `, ``, -1)
		shell = strings.Replace(shell, `"`, `'`, -1)
		shell = strings.Replace(shell, ":'", `:'\`, -1)
		shell = strings.Replace(shell, "$group", `\$group`, -1)
		shell = strings.Replace(shell, "$first", `\$first`, -1)
	}

	return shell
}

//isShowDbs.
func (m Mgo) isShowDbs(shell string) bool {
	if strings.Contains(shell, "showdbs") {
		return true
	}
	return false
}

//isShowCollections.
func (m Mgo) isShowCollections(shell string) bool {
	if strings.Contains(shell, "showcollections") {
		return true
	}
	return false
}

//isAggregate .
func (m Mgo) isAggregate(shell string) bool {
	if strings.Contains(shell, "aggregate") {
		return true
	}
	return false
}

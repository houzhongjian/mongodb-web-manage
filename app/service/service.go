package service

import (
	"dbclient.net/web/mongodb/app/library"
	"dbclient.net/web/mongodb/app/utils"
)

//Service .
type Service struct {
	Utils utils.Utils
	Srv   library.Srv
	Base  library.Base
	Mgo   library.Mgo
}

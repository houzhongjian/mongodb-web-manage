package service

import (
	"github.com/houzhongjian/mongodb-web-manage/app/library"
	"github.com/houzhongjian/mongodb-web-manage/app/utils"
)

//Service .
type Service struct {
	Utils utils.Utils
	Srv   library.Srv
	Base  library.Base
	Mgo   library.Mgo
}

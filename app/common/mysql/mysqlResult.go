package mysql

import (
	"gorm.io/gorm"
	"temp/app/constants/errorCode"
	util "temp/app/utils"
)

type Result struct {
	Db *gorm.DB
}

func NewResult(Db *gorm.DB) *Result {
	return &Result{Db: Db}
}

func (r *Result) Unwrap() (tx *gorm.DB) {
	if r.Db.Error != nil {
		//fmt.Println(r.Db.Error.Error())
		//return r.Db
		//panic(r.Db.Error.Error())
		util.Fail(errorCode.BAD_REQUEST, r.Db.Error.Error())
	}
	return r.Db
}

func (r *Result) UnwrapOr(code int, str string) (tx *gorm.DB) {
	if r.Db.Error != nil {
		util.Fail(code, str)
	}
	return r.Db
}

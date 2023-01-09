package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"temp/app/entity/dto"
	"temp/app/entity/qo"
	"temp/app/service/tempUserService"
	util "temp/app/utils"
	"temp/app/utils/paramsBindEntity"
)

type ResultToken struct {
	Token string
}

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	loginDto = paramsBindEntity.Bind(c, loginDto)
	fmt.Println("dto:", loginDto)
	resultToken := tempUserService.Login(loginDto)

	util.Success(c, resultToken)
}

func Register(c *gin.Context) {
	var registerDto dto.RegisterDto
	registerDto = paramsBindEntity.Bind(c, registerDto)
	res := tempUserService.Register(registerDto)

	util.Success(c, res)
}

func Edit(c *gin.Context) {
	var editUserDto dto.EditUserDto
	editUserDto = paramsBindEntity.Bind(c, editUserDto)
	if editUserDto.ID == 0 {
		editUserDto.ID = util.GetID(c)
	}
	tempUserService.Edit(editUserDto)

	util.Success(c, true)
}

func Destroy(c *gin.Context) {
	data, _ := c.GetRawData()
	toJson := util.ByteToJson(data)

	tempUserService.Destroy(int64(toJson["id"].(float64)))

	util.Success(c, true)
}

func Paginator(c *gin.Context) {
	var paginatorQo qo.PaginatorQo
	paginatorQo = paramsBindEntity.Bind(c, paginatorQo)

	var searchQo qo.SearchTempUserQo
	searchQo = paramsBindEntity.Bind(c, searchQo)

	paginator := tempUserService.Paginator(paginatorQo, searchQo)

	util.Success(c, paginator)
}

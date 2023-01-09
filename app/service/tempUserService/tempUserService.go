package tempUserService

import (
	"fmt"
	"strconv"
	"temp/app/constants/errorCode"
	"temp/app/entity/dto"
	"temp/app/entity/qo"
	"temp/app/model"
	util "temp/app/utils"
	"temp/app/utils/paginator"
)

type ResultToken struct {
	Token string
}

func Login(loginDto dto.LoginDto) ResultToken {
	tempUser := model.TempUser{Username: loginDto.Username}
	user := util.NewResult(tempUser.FindByUsername()).UnwrapOr(errorCode.LOGIN_ERROR, "根据用户名查询数据失败")
	util.NewResult(util.CheckPassword(loginDto.Password, user.PasswordSalt, user.Password)).UnwrapOr(errorCode.LOGIN_ERROR, "账户或密码错误，请重试")

	token, _ := util.GenerateToken(util.HmacUser{
		Id:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
	})

	return ResultToken{Token: token}
}

func Register(registerDto dto.RegisterDto) dto.RegisterDto {
	tempUser := model.TempUser{Username: registerDto.Username}
	find, _ := tempUser.FindByUsername()
	if find != nil {
		util.Fail(errorCode.BAD_REQUEST, "用户名已存在")
	}

	salt := util.RandStr(60)
	pwd := util.GenPwd(registerDto.Password, salt)
	tempUser.Password = pwd
	tempUser.PasswordSalt = salt

	_, _ = tempUser.Create()

	return registerDto
}

func Edit(editUserDto dto.EditUserDto) bool {
	tempUser := model.TempUser{ID: editUserDto.ID}
	ret := util.NewResult(tempUser.FindFromCache()).Unwrap()

	if editUserDto.Username != "" {
		tempUser.Username = editUserDto.Username
	}
	if editUserDto.Password != "" {
		tempUser.Password = util.GenPwd(editUserDto.Password, ret["password_salt"])
	}
	_ = util.NewResult(tempUser.Update()).Unwrap()

	return true
}

func Destroy(id int64) bool {
	tempUser := model.TempUser{ID: id}
	user, _ := tempUser.Find()
	fmt.Println(user)
	if user.ID == 0 {
		util.Fail(errorCode.BAD_REQUEST, "id不存在")
	}
	_ = util.NewResult(tempUser.Destroy()).Unwrap()

	return true
}

func Paginator(paginatorQo qo.PaginatorQo, searchQo qo.SearchTempUserQo) *paginator.PaginatorCollection[model.TempUser] {
	tempUser := model.TempUser{}
	items, total := tempUser.ItemsAndTotal(paginatorQo, searchQo)
	paginatorCollection := paginator.NewPaginatorCollection[model.TempUser](paginatorQo.Page, *items, paginatorQo.PageSize, total)
	return paginatorCollection
}

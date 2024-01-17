package repository

import (
	"BsmgRefactoring/define"
	"fmt"
)

func (sr structBsmgRepository) SelectUserList() (userList []define.BsmgMemberInfo, err error) {
	fmt.Println("Repository GetUserList")

	userList, err = sr.dm.SelectUserList()

	return userList, err
}

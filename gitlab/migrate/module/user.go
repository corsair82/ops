package module

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"regexp"
	"strconv"
)

func GetUsers(gitlabapi, token string) []User {
	var userarr []User
	usersPages, _ := GetPagestotal(gitlabapi+"/users", token)
	for i := 1; i <= usersPages; i++ {
		body, _ := GetBody(gitlabapi+"/users?page="+strconv.Itoa(i), token)
		var users []User
		json.Unmarshal(body, &users)
		userarr = append(userarr, users...)
	}
	return userarr
}
func CheckeUser(gitlabapi, token string, userinfo User) error {

	dstusers := GetUsers(gitlabapi, token)
	for _, user := range dstusers {
		if userinfo.Username == user.Username {
			return errors.New("用户已存在")
		}
	}
	return nil
}

func NewCreatUserArgs(userinfo User) *CreateUserArgs {
	projectslimit := strconv.Itoa(int(userinfo.Projects_limit))
	isadmin := "false"
	if userinfo.Is_admin {
		isadmin = "true"
	}
	return &CreateUserArgs{
		Username:       userinfo.Username,
		Name:           userinfo.Name,
		Email:          userinfo.Email,
		Admin:          isadmin,
		Projects_limit: projectslimit,
		Provider:       userinfo.Identities[0].Provider,
		Extern_uid:     userinfo.Identities[0].Extern_uid,
	}
}

func CreateUser(gitlabapi, token string, userinfo User) error {
	matched, _ := regexp.MatchString(`(root|(.+-boot))`, userinfo.Username)
	if matched {
		return errors.New("系统内置账号，不需要生成！！！")
	}
	url := gitlabapi + "/users"
	userargs := NewCreatUserArgs(userinfo)
	usercheck := CheckeUser(gitlabapi, token, userinfo)
	if usercheck != nil {
		log.Println(usercheck)
	}
	postuserargs, _ := json.Marshal(userargs)
	resp, err := PostHttp(url, token, bytes.NewBuffer(postuserargs))
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	log.Println("RequestUrl:", url, *userargs)
	log.Println(string(body))
	log.Println("The user", userinfo.Username, ":", resp.Status)
	return nil
}

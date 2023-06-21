package module

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"strconv"
)

func GetGroups(gitlabapi, token string) []ProjectGroups {
	var groupsarr []ProjectGroups
	groupsurl := gitlabapi + "/groups"
	groupsPages, _ := GetPagestotal(groupsurl, token)
	for i := 1; i <= groupsPages; i++ {
		body, _ := GetBody(groupsurl+"?page="+strconv.Itoa(i), token)
		var group []ProjectGroups
		json.Unmarshal(body, &group)
		groupsarr = append(groupsarr, group...)
	}
	return groupsarr

}

func GetGroupContainProjects(gitlabapi, token string) (map[string][]Project, error) {
	groups := GetGroups(gitlabapi, token)
	groupcontainprojects := make(map[string][]Project)
	for _, v := range groups {
		groupid := int(v.Id)
		url := gitlabapi + "/groups/" + strconv.Itoa(groupid)
		body, err := GetBody(url, token)
		if err != nil {
			log.Println("Access"+url+"Response Error:", err)
			return nil, err
		}
		var group ProjectGroup
		json.Unmarshal(body, &group)
		groupcontainprojects[group.Name] = group.Projects
	}
	return groupcontainprojects, nil
}

func CreateGroup(gitlabapi, token string, groupinfo ProjectGroups) error {
	url := gitlabapi + "/groups"
	groupargs := NewCreatGroupArgs(groupinfo)

	postgroupargs, _ := json.Marshal(groupargs)
	resp, err := PostHttp(url, token, bytes.NewBuffer(postgroupargs))
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	log.Println("RequestUrl:", url, *groupargs)
	log.Println(string(body))
	log.Println("Create the Group ", groupargs.Name, ":", resp.Status)
	return nil
}

func NewCreatGroupArgs(groupinfo ProjectGroups) *CreateGroupArg {
	return &CreateGroupArg{
		Name:        groupinfo.Name,
		Description: groupinfo.Description,
		Path:        groupinfo.Path,
		Visibility:  groupinfo.Visibility,
	}
}

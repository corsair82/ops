package module

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"strconv"
	"strings"
)

func GetProjects(gitlabapi, token string) []Project {
	var projectsarr []Project
	apiinfoarr := strings.Split(gitlabapi, "/")
	apiversion := apiinfoarr[len(apiinfoarr)-1]
	var url string
	if apiversion == "v3" {
		url = gitlabapi + "/projects/all"
	} else {
		url = gitlabapi + "/projects"
	}

	projectsPages, _ := GetPagestotal(url, token)
	for i := 1; i <= projectsPages; i++ {
		//body, _ := GetBody(gitlabapi+"/projects?page="+strconv.Itoa(i), token)
		body, _ := GetBody(url+"?page="+strconv.Itoa(i), token)
		var project []Project
		json.Unmarshal(body, &project)
		projectsarr = append(projectsarr, project...)
	}
	return projectsarr

}

func GetProjectsMembers(gitlabapi, token string) (map[string][]Member, error) {
	projecttomembermap := make(map[string][]Member)
	projectsarr := GetProjects(gitlabapi, token)
	for _, v := range projectsarr {
		projectid := int(v.Id)
		url := gitlabapi + "/projects/" + strconv.Itoa(projectid) + "/members/all"
		body, err := GetBody(url, token)
		if err != nil {
			log.Println("Access"+url+"Response Error:", err)
			return nil, err
		}
		var members []Member
		json.Unmarshal(body, &members)
		projecttomembermap[v.Name] = members

	}
	return projecttomembermap, nil
}

func CreateProject(gitlabapi, token string, projectinfo Project, groups []ProjectGroups) (Project, error) {
	url := gitlabapi + "/projects"
	var namespaceid float32
	var namespaceidstr string
	newproject := Project{}
	for _, g := range groups {
		if projectinfo.Namespace.Name == g.Name {
			namespaceid = g.Id
			break
		}
	}
	if namespaceid != 0 {
		namespaceidstr = strconv.Itoa(int(namespaceid))
	} else {
		namespaceidstr = ""
	}
	createProjectArg := NewCreatProjectArg(projectinfo, namespaceidstr)
	postprojectargs, _ := json.Marshal(createProjectArg)
	resp, err := PostHttp(url, token, bytes.NewBuffer(postprojectargs))
	if err != nil {
		return newproject, err
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &newproject)
/*
	log.Println("RequestUrl:", url, createProjectArg)
	log.Println("ProjectId:", newproject.Id, "ProjectName:", newproject.Name)
	log.Println("The ProjectName", projectinfo.Name, ":", resp.Status)
	log.Println("The Body:", string(body))
*/
	return newproject, nil

}

/*
	func CreateProjectToUser(gitlabapi, token string, creatorid float32, projectinfo Project) error {
		url := gitlabapi + "/projects/user/" + strconv.Itoa(int(creatorid))
		createprojecttouserarg := NewCreatProjectArg(projectinfo)
		postprojecttouserargs, _ := json.Marshal(createprojecttouserarg)
		resp, err := PostHttp(url, token, bytes.NewBuffer(postprojecttouserargs))
		if err != nil {
			return err
		}
		body, _ := io.ReadAll(resp.Body)
		log.Println("RequestUrl:", url, createprojecttouserarg)
		log.Println(string(body))
		log.Println("The ProjectName", projectinfo.Name, ":", resp.Status)
		return nil
	}
*/
func NewCreatProjectArg(projectinfo Project, namespaceid string) *CreateProjectArg {

	return &CreateProjectArg{
		Name:               projectinfo.Name,
		Description:        projectinfo.Description,
		Merge_method:       projectinfo.Merge_method,
		Visibility:         projectinfo.Visibility,
		Pages_access_level: projectinfo.Pages_access_level,
		Namespace_id:       namespaceid,
		Tag_list:           projectinfo.Tag_list,
	}

}

func JoinUserToProject(gitlabapi, token string, p Project, m []Member) error {
	url := gitlabapi + "/projects/" + strconv.Itoa(int(p.Id)) + "/members"
	postprojectmemberags, _ := json.Marshal(m)
	resp, err := PostHttp(url, token, bytes.NewBuffer(postprojectmemberags))
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	log.Println("RequestUrl:", url, m)
	log.Println(string(body))
	log.Println("The ProjectName", p.Name, "jion these users", m, ":", resp.Status)
	return nil
}

func NewCreateMember(m Member, user User) *Member {
	return &Member{
		Id:           user.Id,
		Name:         m.Name,
		UserName:     m.UserName,
		Access_Level: m.Access_Level,
		State:        m.State,
	}

}

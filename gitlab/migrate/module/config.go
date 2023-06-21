package module

type User struct {
	Id                 float32    `json:"id"`
	Name               string     `json:"name"`
	Username           string     `json:"username"`
	State              string     `json:"state"`
	Email              string     `json:"email"`
	Can_create_group   bool       `json:"can_create_group"`
	Can_create_project bool       `json:"can_create_project"`
	Projects_limit     float32    `json:"projects_limit"`
	Is_admin           bool       `json:"is_admin"`
	Identities         []Identity `json:"identities"`
}
type Identity struct {
	Provider   string `json:"provider"`
	Extern_uid string `json:"extern_uid"`
}
type CreateUserArgs struct {
	Password         string `json:"password,omitempty"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Admin            string `json:"admin"`
	Can_create_group bool   `json:"can_create_group,omitempty"`
	Projects_limit   string `json:"projects_limit,omitempty"`
	Provider         string `json:"provider,omitempty"`
	Extern_uid       string `json:"extern_uid,omitempty"`
}
type Namespace struct {
	Id        float32 `json:"id"`
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	Kind      string  `json:"kind"`
	Full_path string  `json:"full_path"`
	Parent_id float32 `json:"parent_id"`
}
type Share_Groups struct {
	Group_id           float32 `json:"group_id"`
	Group_name         string  `json:"group_name"`
	Group_path         string  `json:"group_path"`
	Group_access_level float32 `json:"group_access_leve"`
}
type Project struct {
	Id                  float32        `json:"id"`
	Name                string         `json:"name"`
	Name_with_namespace string         `json:"name_with_namepace,omitempty"`
	Namespace           Namespace      `json:"namespace"`
	Description         string         `json:"description"`
	Visibility          string         `json:"visibility"`
	Runners_token       string         `json:"runners_token"`
	Share_with_groups   []Share_Groups `json:"share_with_groups"`
	Creator_id          float32        `json:"creator_id"` //创建用户是需要调用/projects/user/:user_id
	Merge_method        string         `json:"merge_method"`
	Pages_access_level  string         `json:"pages_access_level"`
	Http_url_to_repo    string         `json:"http_url_to_repo"`
	Ssh_url_to_repo     string         `json:"ssh_url_to_repo"`
	Tag_list            []string       `json:"tag_list"`
}

type CreateProjectArg struct {
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Merge_method       string   `json:"merge_method,omitempty"`
	Visibility         string   `json:"visibility,omitempty"`
	Namespace_id       string   `json:"namespace_id,omitempty"`
	Pages_access_level string   `json:"pages_access_level,omitempty"`
	User_id            string   `json:"user_id,omitempty"` //owner userid
	Tag_list           []string `json:"tag_list"`
}

type Member struct {
	Id           float32 `json:"id"`
	Name         string  `json:"name"`
	UserName     string  `json:"username"`
	Access_Level float32 `json:"access_level"`
	State        string  `json:"state"`
}

type ProjectGroups struct {
	Id                      float32 `json:"id"`
	Name                    string  `json:"name"`
	Path                    string  `json:"path"`
	Visibility              string  `json:"visibility"`
	Full_name               string  `json:"full_name"`
	Full_path               string  `json:"full_path"`
	Description             string  `json:"description"`
	Parent_id               float32 `json:"parent_id"`
	Project_creation_level  float32 `json:"project_creation_level"`
	Subgroup_creation_level float32 `json:"subgroup_creation_level"`
	Ifs_enabled             bool    `json:"ifs_enabled"`
	Request_access_enabled  bool    `json:"request_access_enabled"`
}
type ProjectGroup struct {
	Id                      float32   `json:"id"`
	Name                    string    `json:"name"`
	Visibility              string    `json:"visibility"`
	Full_name               string    `json:"full_name"`
	Full_path               string    `json:"full_path"`
	Description             string    `json:"description"`
	Parent_id               float32   `json:"parent_id"`
	Project_creation_level  float32   `json:"project_creation_level"`
	Subgroup_creation_level float32   `json:"subgroup_creation_level"`
	Ifs_enabled             bool      `json:"ifs_enabled"`
	Request_access_enabled  bool      `json:"request_access_enabled"`
	Runner_token            string    `json:"runner_token"`
	Projects                []Project `json:"projects"`
}
type CreateGroupArg struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Visibility  string `json:"visibility,omitempty"`
}

type PushMirrorArg struct {
	RemoteRepoUrl string `json:"url"`
}

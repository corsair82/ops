package main

const (
	SRC_PRIVATE_TOKEN = ""
	SRC_GITLAB_API_V4 = "https://gitlab-ui.1234.com/api/v4"
	DST_PRIVATE_TOKEN = "glpat-zhRsZGZMS8D_3RDsRyF1"
	DST_GITLAB_API_V4 = "http://10.1.87.200/api/v4"
	NoAccess          = 0
	MinimalAccess     = 5
	Guest             = 10
	Reporter          = 20
	Developer         = 30
	Maintainer        = 40
	Owner             = 50
)

func main() {
	//项目
	//项目组
	/*
		groups := gxmod.GetGroups("https://gitlab-ui.1234.com/api/v4", "8j76rFM7a-DaySLjY1WJ")

		for _, g := range groups {
			err := gxmod.CreateGroup("http://172.22.7.83/api/v4", "DS2UyzgxMoryyswLsGUy", g)
			if err != nil {
				log.Panicln(err)
			}
		}
	*/

	//找项目
	/*
		srcprojectsarr := gxmod.GetProjects("http://gitlab.1234.com/api/v3", "5cNHMkncsQg8QbEziz3v")
		//srcprojectsarr := gxmod.GetProjects("https://gitlab-ui.1234.com/api/v4", "8j76rFM7a-DaySLjY1WJ")
		dstprojectsarr := gxmod.GetProjects("http://172.22.7.83/api/v4", "DS2UyzgxMoryyswLsGUy")

		for _, srcp := range srcprojectsarr {
			var findpro int
			for _, dstp := range dstprojectsarr {

				if srcp.Name == dstp.Name && srcp.Name_with_namespace == dstp.Name_with_namespace {
					findpro = 1
					fmt.Println("FindPro:", srcp.Name, srcp.Id, srcp.Ssh_url_to_repo, dstp.Id, dstp.Ssh_url_to_repo)
					break
				}
				if srcp.Name == dstp.Name {
					findpro = 1
					fmt.Println("FindPro:", srcp.Name, srcp.Id, srcp.Ssh_url_to_repo, dstp.Id, dstp.Ssh_url_to_repo)
					break
				}
			}
			if findpro == 0 {
				fmt.Println("NotFindPro:", srcp.Name, srcp.Id)
			}

		}
	*/
	/*
	   //创建项目
	   dstgroups := gxmod.GetGroups("http://172.22.7.83/api/v4", "DS2UyzgxMoryyswLsGUy")
	   projectsarr := gxmod.GetProjects("https://gitlab-ui.1234.com/api/v4", "8j76rFM7a-DaySLjY1WJ")

	   	for _, p := range projectsarr {
	   		newproject, err := gxmod.CreateProject("http://172.22.7.83/api/v4", "DS2UyzgxMoryyswLsGUy", p, dstgroups)
	   		if err != nil {
	   			log.Panicln(err)
	   		}
	   		fmt.Println(p.Name, p.Id, p.Ssh_url_to_repo, newproject.Id, newproject.Ssh_url_to_repo)

	   }
	*/
}

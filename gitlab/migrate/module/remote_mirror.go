package module

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"strconv"
	"strings"
)

// curl --request POST --header "PRIVATE-TOKEN: p_HzMQVskpRdPYUMyx7w" --data "url=http://fengjj:junr3092#@gitlab-ui.htexam.com/fengjj2/scripts3.git" http://gitlab.huatuop.com/api/v4/projects/506/remote_mirrors
func CreatePushMirror(gitlabapi, srctoken, dsttoken string, srcprojectinfo, dstprojectinfo Project) error {

	url := gitlabapi + "/projects/" + strconv.Itoa(int(srcprojectinfo.Id)) + "/remote_mirrors"
	dsthttpurl := strings.ReplaceAll(dstprojectinfo.Http_url_to_repo, "http://", "")
	mirrorurl := "http://" + "gongxuan:" + dsttoken + "@" + dsthttpurl
	createpushmirrorarg := PushMirrorArg{
		RemoteRepoUrl: mirrorurl,
	}
	pushmirrorargs, _ := json.Marshal(createpushmirrorarg)
	resp, err := PostHttp(url, srctoken, bytes.NewBuffer(pushmirrorargs))
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	log.Println("RequestUrl:", url, createpushmirrorarg)
	log.Println(string(body))

	return nil

}

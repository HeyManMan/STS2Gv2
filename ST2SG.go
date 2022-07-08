package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"ST2G/cvemod/s001"
	"ST2G/cvemod/s005"
	"ST2G/cvemod/s007"
	"ST2G/cvemod/s008"
	"ST2G/cvemod/s009"
	"ST2G/cvemod/s012"
	"ST2G/cvemod/s013"
	"ST2G/cvemod/s015"
	"ST2G/cvemod/s016"
	"ST2G/cvemod/s045"
	"ST2G/cvemod/s046"
	"ST2G/cvemod/s048"
	"ST2G/cvemod/s053"
	"ST2G/cvemod/s057"
	"ST2G/cvemod/utils"

	"github.com/urfave/cli/v2"
)

//todo: 添加多线程处理
//todo: 添加批量扫描
//todo: 优化程序处理逻辑
//todo: 优化日志输出显示
//todo: 优化代码、调整格式

type oneRes struct {
	url       string
	data      string
	cmd       string
	mode      string
	vulNumber int
}

var (
	mode      string
	url       string
	data      string
	cmd       string
	vulNumber int
	thread    int
	filename  string
)

func buildRes(url, data, cmd, mode, filename string, vulNumber int) (res []oneRes) {
	var vulList = [...]int{1, 5, 7, 8, 9, 12, 13, 15, 16, 45, 46, 48, 53, 57}
	var mustHavePostDate = [...]int{1, 7, 12, 48, 53}
	var mustHaveGetDate = [...]int{9}
	var noneParList = [...]int{5, 8, 13, 15, 16, 45, 46, 48, 57}

	if url == "" && filename == "" {
		log.Fatal("请输入URL或者filename")
	}
	if data == "" {
		for _, i := range mustHaveGetDate {
			if i == vulNumber {
				log.Fatalf("Struts2-0%d需指定要测试的GET参数，如: --data=\"name\"", vulNumber)
			}
		}
		for _, i := range mustHavePostDate {
			if i == vulNumber {
				log.Fatalf("Struts2-0%d需指定POST数据包内容，并用<fuckit>标记出测试点，如: --data \"user=a&pass=fuckit\"", vulNumber)
			}
		}
	}
	if mode == "exploit" && cmd == "" {
		log.Fatal("exploit模式必须指定cmd参数")
	}

	flag := false
	for _, i := range vulList {
		if i == vulNumber {
			flag = true
		}
	}

	if filename != "" {
		d, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Panic(err)
		}

		// todo:文件去重
		lines := strings.Split(string(d), "\n")
		if vulNumber == 0 {
			for _, line := range lines {
				for _, i := range noneParList {
					r := oneRes{
						url:       line,
						data:      data,
						cmd:       cmd,
						mode:      mode,
						vulNumber: i,
					}
					res = append(res, r)
				}
			}
		} else if flag {
			for _, line := range lines {
				r := oneRes{
					url:       line,
					data:      data,
					cmd:       cmd,
					mode:      mode,
					vulNumber: vulNumber,
				}
				res = append(res, r)
			}
		} else {
			log.Fatal("vulNumber 错误")
		}
	} else if vulNumber == 0 && filename == "" && url != "" {
		for _, i := range noneParList {
			r := oneRes{
				url:       url,
				data:      data,
				cmd:       cmd,
				mode:      mode,
				vulNumber: i,
			}
			res = append(res, r)
		}
	} else if flag && url != "" {
		r := oneRes{
			url:       url,
			data:      data,
			cmd:       cmd,
			mode:      mode,
			vulNumber: vulNumber,
		}
		res = append(res, r)
	} else {
		log.Fatal("vulNumber 错误")
	}

	return res
}

func run(res chan oneRes, wg *sync.WaitGroup) {
	defer wg.Done()
	r := <-res
	if r.mode == "check" {
		switch r.vulNumber {
		case 1:
			s001.Check(r.url, r.data)
		case 5:
			s005.Check(r.url)
		case 7:
			s007.Check(r.url, r.data)
		case 8:
			s008.Check(r.url)
		case 9:
			s009.Check(r.url, r.data)
		case 12:
			s012.Check(r.url, r.data)
		case 13:
			s013.Check(r.url)
		case 15:
			s015.Check(r.url)
		case 16:
			s016.Check(r.url)
		case 45:
			s045.Check(r.url)
		case 46:
			s046.Check(r.url)
		case 48:
			s048.Check(r.url, r.data)
		case 53:
			s053.Check(r.url, r.data)
		case 57:
			s057.Check(r.url)
		default:
			//todo:list cve
			fmt.Println("漏洞编号设置错误，目前支持检测：")
			for _, vnn := range utils.VulLIst {
				fmt.Println(vnn)
			}
		}
	} else if r.mode == "exploit" {
		switch r.vulNumber {
		case 1:
			s001.ExecCommand(r.url, r.cmd, r.data)
		case 5:
			s005.ExecCommand(r.url, r.cmd)
		case 7:
			s007.ExecCommand(r.url, r.cmd, r.data)
		case 8:
			s008.ExecCommand(r.url, r.cmd)
		case 9:
			s009.ExecCommand(r.url, r.cmd, r.data)
		case 12:
			s012.ExecCommand(r.url, r.cmd, r.data)
		case 13:
			s013.ExecCommand(r.url, r.cmd)
		case 15:
			s015.ExecCommand(r.url, r.cmd)
		case 16:
			s016.ExecCommand(r.url, r.cmd)
		case 45:
			s045.ExecCommand(r.url, r.cmd)
		case 46:
			s046.ExecCommand(r.url, r.cmd)
		case 48:
			s048.ExecCommand(r.url, r.cmd, r.data)
		case 53:
			s053.ExecCommand(r.url, r.cmd, r.data)
		case 57:
			s057.ExecCommand(r.url, r.cmd)
		default:
			log.Fatalf("命令执行模式必须指定漏洞编号")
		}
	}
}

func action(c *cli.Context) error {
	var res []oneRes

	res = buildRes(url, data, cmd, mode, filename, vulNumber)
	var wg sync.WaitGroup
	resC := make(chan oneRes, thread)

	for _, r := range res {
		wg.Add(1)

		resC <- r
		go run(resC, &wg)
	}

	//go func() {
	//	for _, r := range res {
	//		resC <- r
	//	}
	//}()
	//
	//go func() {
	//	wg.Add(1)
	//	run(resC, &wg)
	//}()

	defer close(resC)
	wg.Wait()
	return nil
}

func main() {
	app := &cli.App{
		Name:  "ST2SG",
		Usage: "Struts2 Scanner Written in Golang",
		UsageText: "ST2SG -f target.txt\n" +
			"   ST2SG -u https://test.com/test.action\n" +
			"   ST2SG -u https://test.com/test.action --vn 15 -m check\n",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mode",
				Aliases:     []string{"m"},
				Usage:       "Work mode, check or exploit",
				Value:       "check",
				Destination: &mode,
			},
			&cli.IntFlag{
				Name:        "vul",
				Aliases:     []string{"n"},
				Usage:       "Vulnerability number",
				Value:       0,
				Destination: &vulNumber,
			},
			&cli.StringFlag{
				Name:        "url",
				Aliases:     []string{"u"},
				Usage:       "Set target url",
				Destination: &url,
			},
			&cli.StringFlag{
				Name:        "cmd",
				Aliases:     []string{"c"},
				Usage:       "Exec command(Only works on mode exploit.)",
				Destination: &cmd,
			},
			&cli.StringFlag{
				Name:        "data",
				Aliases:     []string{"d"},
				Usage:       "Specific vulnerability packets",
				Destination: &data,
			},
			&cli.StringFlag{
				Name:        "filename",
				Aliases:     []string{"f"},
				Usage:       "set target url file",
				Destination: &filename,
			},
			&cli.IntFlag{
				Name:        "thread",
				Aliases:     []string{"t"},
				Usage:       "set thread",
				Value:       20,
				Destination: &thread,
			},
		},

		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

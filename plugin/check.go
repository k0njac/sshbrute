package plugin

import (
	"fmt"
	"net"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

var mutex sync.Mutex

func CheckPortAlive(ip_port string) {
	_, err := net.DialTimeout("tcp", ip_port, time.Second*5)
	if err == nil {
		mutex.Lock()
		fmt.Printf("[+]%v is alive\n", ip_port)
		Checkiplists = append(Checkiplists, ip_port)
		mutex.Unlock()
	} else {
		fmt.Printf("[-]%v is no alive\n", ip_port)
	}
}

var Success map[string]bool

func Checkinit() {
	Success = make(map[string]bool)
}

func CheckSsh(usernmae string, password string, ip string, command string) bool {
	if Success[ip] {
		return true
	}
	config := &ssh.ClientConfig{
		User: usernmae,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if Success[ip] {
		return true
	}
	client, err := ssh.Dial("tcp", ip, config)
	if err == nil {
		defer client.Close()
		Success[ip] = true
		fmt.Printf("[+] ssh: %v  %v %v\n", ip, usernmae, password)
		if command != "" {
			session, err := client.NewSession()
			if err != nil {
				fmt.Println("[-] 创建ssh session 失败", err)
			}
			Ret, err := session.CombinedOutput(command)
			if err == nil {
				defer session.Close()
				fmt.Println("[+] 执行命令结果", string(Ret))
			} else {
				fmt.Println("[-] 命令执行失败")
			}
		}
	}
	return true
}

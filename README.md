虽然ssh爆破有很多，但是轮询的还是少，某次项目中对方的waf特别猛，所以单独针对ip爆破总是被封。所以造了个轮询的ssh爆破

## 使用指南

ip.txt 填写ip 或 ip:端口
不加端口则默认使用22
```
1.1.1.1:22
1.1.1.2
1.1.1.3:2222
...
```

user.txt 添加要爆破的用户名
```
root
admin
```

password.txt 则填写要爆破的字典

plugin/getdict.go:89行 
```
            CheckSsh(value.User, value.Pass, value.Ip, "w") // 可以替换成whoami
```
可以替换你默认登录以后要执行的命令 如果不想执行任何命令 则为空
# tel.0sn.net

```bash
$ telnet localhost 8023
Trying ::1...
Connected to localhost.
Escape character is '^]'.

--------------------------
tel.0sn.netへようこそ！
_       _   ___                         _   
| |_ ___| | / _ \ ___ _ __    _ __   ___| |_ 
| __/ _ \ || | | / __| '_ \  | '_ \ / _ \ __|
| ||  __/ || |_| \__ \ | | |_| | | |  __/ |_ 
 \__\___|_(_)___/|___/_| |_(_)_| |_|\___|\__|
--------------------------
あなたは 1 人目の訪問者です。
--------------------------
Web: https://0sn.net
--------------------------
Connection closed by foreign host.
```

## 使い方
```bash
$ go build -o main

#デフォルトでは8023で待ち受け
$ ./main
$ PORT=8000 ./main
```

### Docker
```bash
$ docker compose build

$ docker compose up -d
```
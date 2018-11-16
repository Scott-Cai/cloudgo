package main

import (
	"fmt"
	"os"
	"log"

	"github.com/Scott-Cai/cloudgo/service"
	flag "github.com/spf13/pflag"
)

const (
	// PORT : default port
	PORT string = "8080"
)

var port string

func init()  {
	// ���������в���
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	// ���������в���
	flag.Parse()

	// ���������˿�
	// �û����õ�һ��ϵͳ������֮��Ĭ�����
	if len(*pPort) != 0 {
		port = *pPort
	} else {
		if port = os.Getenv("PORT"); len(port) == 0 {
			port = PORT
		}
	}
}

func main() {
	// ��ʼ����
	if err := service.Run(port); err != nil {
        log.Fatal("start fasthttp fail:", err)
    } else {
		fmt.Println("listening on :" + port)
	}
}

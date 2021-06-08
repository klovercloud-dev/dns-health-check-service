package helper

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"unicode"
)

func Curl(url string) bool {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("inside false: ", url)
		return false
	}
	if res.StatusCode == 200 {
		return true
	}
	return false
}

func TcpConnect(ips []string, port string) bool {
	for _, ip := range ips {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), timeout)
		if err != nil {
			fmt.Println("Connecting error:", err)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", net.JoinHostPort(ip, port))
			return true
		}
	}
	return false
}

func GetRecordType(val string) string {
	var recordType string
	curlyCount := 0
	for _, char := range val {
		if char == '{' {
			curlyCount += 1
		}
		if curlyCount == 2 {
			break
		}
		if unicode.IsLetter(char) {
			recordType += string(char)
		}
	}
	return recordType
}

func Log(title string,val string){
	fmt.Println("-----------------------------------------")
	fmt.Println(title)
	fmt.Println(val)
	fmt.Println("-----------------------------------------")
}


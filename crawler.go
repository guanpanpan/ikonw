// crawler网络爬虫
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func writeToFile() int {
	response, err := http.Get("http://www.cnblogs.com/guanpanpan/") //下划线为空标识符号
	if err != nil {                                                 //未设置时为nil
		log.Fatal(err)
		os.Exit(1)
	}
	defer response.Body.Close() //在函数返回时执行

	buf := make([]byte, 1024)
	blogFile, err1 := os.OpenFile("d:/cnblog.html", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err1 != nil {
		panic(err1)
	}
	defer blogFile.Close()
	for {
		n, _ := response.Body.Read(buf)
		if 0 == n {
			break
		}
		blogFile.WriteString(string(buf[:n]))
	}

	return 0
}

func analyzeDownUrl() []string {
	response, _ := http.Get("http://www.cnblogs.com/guanpanpan/") //下划线为空标识符号
	defer response.Body.Close()                                   //在函数返回时执行
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println("Hello World!")
	str := make([]string, 10)
	return str
}
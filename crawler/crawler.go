// crawler网络爬虫
package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func WriteToFile() int {

	saveFile := os.Getenv("WEBPATH") + "/cnblog.html"
	response, err := http.Get("http://www.cnblogs.com/guanpanpan/") //下划线为空标识符号
	if err != nil {                                                 //未设置时为nil
		log.Fatal(err)
		os.Exit(1)
	}
	defer response.Body.Close() //在函数返回时执行

	buf := make([]byte, 1024)
	blogFile, err1 := os.OpenFile(saveFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err1 != nil {
		panic(err1)
	}
	defer blogFile.Close()
	blogFile.WriteString("update for" + time.Now().String() + "<br>")
	for {
		n, _ := response.Body.Read(buf)
		if 0 == n {
			break
		}
		blogFile.WriteString(string(buf[:n]))
	}

	return 0
}

func AnalyzeDownUrl() []string {
	response, _ := http.Get("http://www.cnblogs.com/guanpanpan/") //下划线为空标识符号
	defer response.Body.Close()                                   //在函数返回时执行
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println("Hello World!")
	str := make([]string, 10)
	return str
}

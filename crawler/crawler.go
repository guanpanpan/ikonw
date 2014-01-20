// crawler网络爬虫
package crawler

import (
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

type iurl struct {
	href  string
	title string
	path  string
}

//写入主页链接
func WriteIndexhtml(iurls []*iurl) {
	saveFile := os.Getenv("WEBPATH") + "/index.html"
	blogFile, _ := os.OpenFile(saveFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, os.ModePerm)
	defer blogFile.Close()
	blogFile.WriteString("<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/>")
	blogFile.WriteString("update for" + time.Now().String() + "<br>")
	for i := range iurls {
		blogFile.WriteString("<a href='" + iurls[i].path + ".html'>" + iurls[i].title + "</a><br>\n")
	}

}

//抓取网页
func (url *iurl) WriteToFile() int {
	saveFile := os.Getenv("WEBPATH") + "/" + url.path + ".html"
	urlStr := getUrlString(url.href)
	blogFile, err1 := os.OpenFile(saveFile, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, os.ModePerm)
	if err1 != nil {
		panic(err1)
	}
	defer blogFile.Close()
	blogFile.WriteString("update for" + time.Now().String() + "<br>")
	blogFile.WriteString(urlStr)
	return 0
}
func getUrlString(url string) string {
	response, _ := http.Get(url) //下划线为空标识符号
	defer response.Body.Close()  //在函数返回时执行
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

//分析要下载网址
func AnalyzeDownUrl(s string, path int) []*iurl {
	urlStr := getUrlString(s)
	//fmt.Println(urlStr)
	//\\S所有字符
	urlRx := regexp.MustCompile("href=\"(\\S*guanpanpan\\S*html)\">(.*)</a>")
	urls := urlRx.FindAllStringSubmatch(urlStr, 100)
	//iurls := make([]*iurl, len(urls))
	//slice
	iurls := []*iurl{}
	for i := range urls {

		if le := len(urls[i]); le == 3 {
			url := iurl{urls[i][1], urls[i][2], fmt.Sprint(path)}
			//iurls[i] = &url
			iurls = append(iurls, &url)

		} else if le == 2 {
			url := iurl{urls[i][1], "empty title", fmt.Sprint(path)}
			//iurls[i] = &url
			iurls = append(iurls, &url)
		}
		path = path + 1
	}

	nextUrl := AnalyzeNextPage(urlStr)
	fmt.Println("next:" + nextUrl)
	if nextUrl != "" {
		nextDownIurls := AnalyzeDownUrl(nextUrl, path)
		iurls = append(iurls, nextDownIurls...)
	}
	return iurls
}

//分析关联页面
func AnalyzeNextPage(s string) string {
	urlRx := regexp.MustCompile("href=\"(\\S*guanpanpan\\S*html\\S*)\">下一页</a>")
	urls := urlRx.FindStringSubmatch(s)
	if len(urls) > 1 {
		return urls[1]
	} else {
		return ""
	}

}
func Down() {
	//分析要下载的url
	iurls := AnalyzeDownUrl("http://www.cnblogs.com/guanpanpan/", 1)
	//写主页
	WriteIndexhtml(iurls)
	//下载网页
	for i := range iurls {
		iurls[i].WriteToFile()
		time.Sleep(time.Second * 10)
	}
	fmt.Println("WriteToFile!%d" + time.Now().String())

}

//抓取网页
//func WriteToFile() int {
//	saveFile := os.Getenv("WEBPATH") + "/cnblog.html"
//	response, err := http.Get("http://www.cnblogs.com/guanpanpan/") //下划线为空标识符号
//	if err != nil {                                                 //未设置时为nil
//		log.Fatal(err)
//		os.Exit(1)
//	}
//	defer response.Body.Close() //在函数返回时执行

//	buf := make([]byte, 1024)
//	blogFile, err1 := os.OpenFile(saveFile, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, os.ModePerm)
//	if err1 != nil {
//		panic(err1)
//	}
//	defer blogFile.Close()
//	blogFile.WriteString("update for" + time.Now().String() + "<br>")
//	for {
//		n, _ := response.Body.Read(buf)
//		if 0 == n {
//			break
//		}
//		blogFile.WriteString(string(buf[:n]))
//	}

//	return 0
//}

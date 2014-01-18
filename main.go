// BackBlog project main.go
package main

import (
	"fmt"
	"ikonw/crawler"
	"ikonw/myserver"
	"time"
)

func main() {
	go func() {
		for {
			crawler.WriteToFile()
			fmt.Println("WriteToFile!%d" + time.Now().String())
			time.Sleep(time.Minute * 30)
		}

	}()

	//crawler.AnalyzeDownUrl()

	myserver.RunMyserver()
	fmt.Println("end!")
}

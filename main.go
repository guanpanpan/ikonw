// BackBlog project main.go
package main

import (
	"fmt"
	"ikonw/crawler"
	"ikonw/myserver"
)

func main() {
	crawler.WriteToFile()
	//crawler.AnalyzeDownUrl()

	myserver.RunMyserver()
	fmt.Println("end!")
}

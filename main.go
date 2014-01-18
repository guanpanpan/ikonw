// BackBlog project main.go
package main

import (
	"fmt"
	"ikonw/crawler"
)

func main() {
	//writeToFile()
	crawler.AnalyzeDownUrl()

	RunMyserver()
	fmt.Println("end!")
}

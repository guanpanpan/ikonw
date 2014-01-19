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
			crawler.Down()
			time.Sleep(time.Minute * 30)
		}

	}()

	myserver.RunMyserver()
	fmt.Println("end!")
}

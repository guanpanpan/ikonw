// myServer
package main

import (
	"github.com/codegangsta/martini"
)

func RunMyserver() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!panpanwangjing"
	})
	m.Get("/test/**", func() string {
		return "Hello world!panpanwangjing test"
	})
	m.Run()
}

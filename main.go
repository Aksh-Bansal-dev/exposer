package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os/exec"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	go func() {
		// cmd := exec.Command("ngrok", "tcp", "8000")
		cmd := exec.Command("echo","http.server")
        err:= cmd.Run()
        if err!=nil{
            log.Fatal(err)
        }
	}()
	time.Sleep(time.Second * 1)
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Println( e.Attr("href"))
		if e.Attr("href")[:2] == "tcp" {
			fmt.Println("found!")
		}
	})
	c.OnHTML("div", func(e *colly.HTMLElement) {
		fmt.Println("found something")
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
    res, err:=http.Get("http://localhost:4040/inspect/http")
    if err!=nil{
        log.Fatal(err)
    }
    dump,_:=httputil.DumpResponse(res,true)
    fmt.Println(string(dump))
    c.Visit("http://localhost:4040/inspect/http")
	fmt.Println("Done")
}

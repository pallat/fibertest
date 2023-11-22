package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go ClientRequest("8080")
	ClientRequest("9090")
	wg.Wait()
}

func ClientRequest(port string) {
	defer wg.Done()
	c := &http.Client{
		Timeout: time.Second * 1,
	}

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:"+port+"/hello", nil)
	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("call successfully")
}

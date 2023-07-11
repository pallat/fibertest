package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

var wg sync.WaitGroup

func main() {
	go serverFiber()
	client()
	wg.Wait()
}

func client() {
	wg.Add(5)
	http.Get("http://127.0.0.1:3000/aaa")
	http.Get("http://127.0.0.1:3000/bbbb")
	http.Get("http://127.0.0.1:3000/c")
	http.Get("http://127.0.0.1:3000/dd")
	http.Get("http://127.0.0.1:3000/eee")
}

func serverFiber() {
	app := fiber.New()
	app.Get("/:test", func(c *fiber.Ctx) error {
		result := c.Params("test")
		go outside(result)
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}

func serverGin() {
	app := gin.New()
	app.GET("/:test", func(c *gin.Context) {
		result := c.Param("test")
		go outside(result)
		c.String(200, "Hello, World!")
	})
	app.Run(":3000")
}

func outside(s string) {
	time.Sleep(time.Millisecond)
	fmt.Printf(">> %d, %v\n", unsafe.StringData(s), s)
	wg.Done()
}

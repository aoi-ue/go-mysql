package main

import (
	"db/infrastructure"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		infrastructure.LoadEnv()
		infrastructure.NewDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	router.Run(":8000")

    sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

	fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")


}

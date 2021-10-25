package main

import (

	"gin/function"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

func main() {

	var wg sync.WaitGroup
	var lock sync.RWMutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	log.Println(Balance(&lock))

	router := gin.Default()
	router.GET("/", function.Saludar1())

	router.GET("/user/:name", function.SaludarFunc())

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", function.ActionTest())

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", function.ActionPost())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(function.CustomRecovery()))

	router.Use(gin.Logger())

	// Per route middleware, you can add as many as you desire.
	router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	router.PUT("/somePut", putting)
	var e = router.Run()

	if e != nil {
		println(e.Error())
	}

}

func benchEndpoint(context *gin.Context) {

}

func MyBenchLogger() gin.HandlerFunc {

	return nil
}

func putting(context *gin.Context) {

}

package main

import (
	"log"
	"quotes-api/internal/infraestructure/router"
	"quotes-api/internal/util/cache"
	"sync"
)

var once sync.Once

func main() {
	go initializeCache()
	router.StartApp()
}

func initializeCache() {
	err := cache.NewBigCacheWrapper()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

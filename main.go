package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

type Example struct {
	Name string
	age  int64
}

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes

	newExample := &Example{
		Name: "testName1",
		age:  22,
	}
	c := cache.New(5*time.Minute, 10*time.Minute)
	foo, found := c.Get("foo")
	if found {
		a := foo.([]Example)
		a = append(a, *newExample)
		c.Set("foo", a, cache.NoExpiration)
		fmt.Println("stored from if part")
	} else {
		c.Set("foo", newExample, cache.NoExpiration)
		fmt.Println("stored from else part")

	}

	foo, found = c.Get("foo")
	if found {
		fmt.Println(foo)
	}

}

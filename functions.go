package main

import "fmt"


func HandlErr(err error, doPanic bool) () {
	fmt.Print("error occured: ", err)
	if doPanic {
		panic(err)
	}
}
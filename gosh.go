package main

import "fmt"

type gosh struct {
}

func (g *gosh) new(args []string) error {
	if newShell, err := ms.createShell(args[0]); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		ms.changeActiveShell(newShell)
	}
	return nil
}

func (g *gosh) set(){
	
}
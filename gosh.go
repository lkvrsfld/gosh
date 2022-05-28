package main

// import (
// 	"errors"
// 	"fmt"
// )

// type gosh struct {
// }

// func (g *gosh) new(args []string) error {
// 	s, err := ms.createShell(args[0])
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	err = g.set(s.name)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (g *gosh) set(name string) error {

// 	s, err := ms.getShellByName(name)

// 	if err != nil {
// 		return errors.New("no shell with name " + name + ", err: " + err.Error())
// 	}
	
// 	err = ms.setInactiveShell(ms.getActiveShell())
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	err = ms.setActiveShell(s)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	return nil
// }

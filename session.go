package main

import (
	"fmt"
	"os"
)

type Session struct {
	shells      []Shell
	activeShell *Shell
    inactiveShells []*Shell
}

func (session *Session) init() {
	newShell, err := session.createShell("gosh1")
	if err != nil {
		fmt.Print(err.Error())
	}
	session.activeShell = newShell
}

func (session *Session) createShell(name string) (*Shell, error) {
    if session.activeShell != nil { 
        fmt.Println("asfasf")
        session.inactiveShells =  append(session.inactiveShells, session.activeShell)
     }
	var newShell Shell
    newShell.name = name
	newShell.id = len(session.shells)
	
	session.shells = append(session.shells, newShell)

	return &newShell, nil
}

func (session *Session) destryShell(s *Shell) error {
    if len(session.inactiveShells) != 0 {
        session.activeShell = session.inactiveShells[len(session.inactiveShells)-1]
        session.inactiveShells = session.inactiveShells[:len(session.inactiveShells)-1]
        session.shells = append(session.shells[:s.id], session.shells[s.id+1:]...)
        return nil
    }
    
    session.exitSession()
    return nil
}

func (session *Session) changeActiveShell(shell *Shell) error {

	session.clearTUI()

	session.activeShell = shell

	return nil

}

func (session *Session) clearTUI() {

}

func (session *Session) exitSession() {
    os.Exit(0)
}

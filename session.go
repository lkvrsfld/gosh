package main

import (
	"errors"
	"fmt"
	"os"
)

type Session struct {
	shells         []Shell
	activeShell    *Shell
	inactiveShells []*Shell
}

func (session *Session) init() {
	newShell, err := session.createShell("gosh1")
	if err != nil {
		fmt.Print(err.Error())
	}
	session.activeShell = newShell
}

func (session *Session) getActiveShell() *Shell {
	return ms.activeShell
}

func (session *Session) getLastInactiveShell() (*Shell, error) {
	if len(session.inactiveShells) != 0 {

		return session.inactiveShells[len(session.inactiveShells)-1], nil
	}
	return nil, errors.New("no last active shell")
}

func (session *Session) getShellByName(name string) (*Shell,error) {
	for _ , shell := range session.shells {
		if shell.name == name {
			return &shell, nil
		}
	}
	return nil, errors.New("shell not found")
}

func (session *Session) createShell(name string) (*Shell, error) {
	if session.activeShell != nil {
		session.inactiveShells = append(session.inactiveShells, session.activeShell)
	}
	var newShell Shell
	newShell.name = name
	newShell.id = len(session.shells)

	session.shells = append(session.shells, newShell)

	return &newShell, nil
}

func (session *Session) destoryShell(s *Shell) error {
	var err error
	// iterate though shells and remove
	if len(session.shells) != 0 {
		for i, shell := range session.shells {
			if &shell == s {
				session.shells = append(session.shells[:i], session.shells[i+1:]...)
				err = nil
				break
			}
			err = errors.New("shell not found")
		}
	}

	session.destoryInactiveShell(s)
	// iterate though inactiveShells and remove
	

	return err
}


func (session *Session) destoryInactiveShell(s *Shell) error {
	var err error
	if len(session.inactiveShells) != 0 {
		for i, inactiveShell := range session.inactiveShells {
			if *inactiveShell == *s {
				session.inactiveShells = append(session.inactiveShells[:i], session.inactiveShells[i+1:]...)
				err = nil
				break
			}
			err = errors.New("inactiveShell not found")
		}
	}
	return err
}



func (session *Session) setActiveShell(s *Shell) error {

	session.clearTUI()
	// onlz needed for new

	err := session.destoryInactiveShell(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	session.activeShell = s
	return nil
}



func (session *Session) setInactiveShell(s *Shell) error {

	// here we need to search session.inactiveShells for shell
	// and if shell already exists, pish it to the latest entry
	// else just add it
	if len(session.inactiveShells) != 0 {
		for i, inactiveShell := range session.inactiveShells {
			fmt.Println("inactiveShell.id")
			fmt.Println(inactiveShell.id)
			fmt.Println("s.id")
			fmt.Println(s.id)
			if *inactiveShell == *s {
				// remove shell
				session.inactiveShells = append(session.inactiveShells[:i], session.inactiveShells[i+1:]...)
				// add shell to end
				session.inactiveShells = append(session.inactiveShells, s)
				return nil
			}

		}
	}
	// add shell to end
	session.inactiveShells = append(session.inactiveShells, s)
	return nil
}

func (session *Session) clearTUI() {

}

func (session *Session) exitSession() {
	os.Exit(0)
}

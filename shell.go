package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
    id int
    name string

    Stdout io.Writer
    Stderr io.Writer
    //history
    //env
    //etc
}


func (s *Shell) start() {

}

func (s *Shell)mainHandler(stdin *bufio.Reader, activeShell *Shell) {

    fmt.Print(s.ps1()) 
    input, err := stdin.ReadString('\n')
    if err != nil {
        fmt.Fprintln(s.Stderr, err)
    }

    // Handle the execution of the input.
    err = s.run(input)
    if err != nil {
        fmt.Fprintln(s.Stderr, err)
    }
    
}

func (s *Shell)run(input string) error {

    // Remove the newline character.
    input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")


    // Prepare the command to execute.
    switch args[0] {
        case "exit":
            s.exit()
        case "gosh":
            if err := s.goshHandler(args[1:]); err != nil {return err}
        default:
            if err := s.shExec(args); err != nil {
                fmt.Println(err.Error())
            }
    }

    // Execute the command and return the error.
    return nil
}


func  (s *Shell)shExec(input []string) error {
    cmd := exec.Command(input[0], input[1:]...)
    cmd.Stderr = s.StdErr()
    cmd.Stdout = s.StdOut()
    return cmd.Run()

}

func (s *Shell) goshHandler(args []string) error {
    var gosh gosh

    switch args[0] {
    case "new":
        if err := gosh.new(args[1:]); err != nil {
            return err
        }
        fmt.Println("created session")
    case "set":
        if err := gosh.set(args[1]); err != nil {
            return err
        }
    case "list":

    }
    return nil
}


func (s *Shell) exit() {
    lastShell, err := ms.getLastInactiveShell()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(0)
    }
    ms.setActiveShell(lastShell)

    ms.destoryShell(s)
}

func (s *Shell) ps1() string {
    return s.name + "> "
}

func (s *Shell) StdErr() io.Writer {
    return s.Stderr
}
func (s *Shell) StdOut() io.Writer {
    return s.Stdout
}


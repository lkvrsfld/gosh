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

    StdOut io.Writer
    StdErr io.Writer
    StdIn *bufio.Reader
    //current dir
    //history

    //env
    //etc
}

func (s *Shell) init(stdin *bufio.Reader, stdout *os.File, stderr *os.File, name string) error {
    s.name = name
    s.id = 0
    s.StdIn = stdin
    s.StdOut = stdout
    s.StdErr = stderr
    return nil
}

// func (s *Shell) start() {

// }

func (s *Shell)mainHandler() {

    fmt.Print(s.ps1()) // change witt line() or something


    // handlesignals

    // handle
    input, err := s.StdIn.ReadString('\n')
    if err != nil {
        fmt.Fprintln(s.StdErr, err)
    }

    // Handle the execution of the input.
    err = s.run(input)
    if err != nil {
        fmt.Fprintln(s.StdErr, err)
    }
    
}

func (s *Shell)run(input string) error {

    // Remove the newline character.
    input = strings.TrimSuffix(input, "\n")
	args := strings.Fields(input)


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
    cmd.Stderr = s.stderr()
    cmd.Stdout = s.stdout()
    return cmd.Run()

}

func (s *Shell) goshHandler(args []string) error {
    // var gosh gosh

    // switch args[0] {
    // case "new":
    //     if err := gosh.new(args[1:]); err != nil {
    //         return err
    //     }
    //     fmt.Println("created session")
    // case "set":
    //     if err := gosh.set(args[1]); err != nil {
    //         return err
    //     }
    // case "list":

    // }
     return nil
}


func (s *Shell) exit() {
    // lastShell, err := ms.getLastInactiveShell()
    // if err != nil {
    //     fmt.Println(err.Error())
    os.Exit(0)
    //}
    // ms.setActiveShell(lastShell)

    // ms.destoryShell(s)
}

func (s *Shell) ps1() string {
    return s.name + "> "
}

func (s *Shell) stderr() io.Writer {
    return s.StdErr
}
func (s *Shell) stdout() io.Writer {
    return s.StdOut
}


package main

import (
	"bufio"
	"os"
)



var ms Session

func main() {
    stdin := stdinListener()
    ms.init()

    
    for {
        activeShell := *ms.activeShell
        activeShell.Stderr = os.Stderr
        activeShell.Stdout = os.Stdout
        activeShell.mainHandler(stdin, ms.activeShell)
    }
}


func stdinListener() *bufio.Reader{
    return bufio.NewReader(os.Stdin)
}


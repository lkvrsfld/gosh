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
        ms.activeShell.Stderr = os.Stderr
        ms.activeShell.Stdout = os.Stdout
        ms.activeShell.mainHandler(stdin, ms.activeShell)
    }
}


func stdinListener() *bufio.Reader{
    return bufio.NewReader(os.Stdin)
}


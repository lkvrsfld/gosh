package main

import (
	"bufio"
	"os"
)

var s Shell

func main() {
    stdin := stdinScanner()
    s.init(stdin, os.Stdout, os.Stderr, "gosh1")
    
    for {
        s.mainHandler()
    }
}

func stdinScanner() *bufio.Reader{
    return bufio.NewReader(os.Stdin)
}

//old using session

// var ms Session

// func main() {
//     stdin := stdinListener()
//     ms.init()

    
//     for {
//         activeShell := *ms.activeShell
//         activeShell.Stderr = os.Stderr
//         activeShell.Stdout = os.Stdout
//         activeShell.mainHandler(stdin, ms.activeShell)
//     }
// }
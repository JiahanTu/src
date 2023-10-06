// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"os/exec"

// 	"github.com/creack/pty"
// )

// func main() {
// 	// Create arbitrary command.
// 	c := exec.Command("./custom_zsh")

// 	// Start the command with a pty.
// 	ptmx, err := pty.Start(c)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Make sure to close the pty at the end.
// 	defer func() { _ = ptmx.Close() }() // Best effort.

// 	// Copy stdin to the pty and the pty to stdout.
// 	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
// 	_, _ = io.Copy(os.Stdout, ptmx)

// 	// Open the recorded history file and output it.
// 	f, err := os.Open("session.log")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	fmt.Println("\nRecorded commands:")
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"log"

	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New("> ")
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF, io.ErrUnexpectedEOF
			break
		}
		log.Printf("你输入了: %s", line)
	}
}

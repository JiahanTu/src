package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/chzyer/readline"
	"github.com/creack/pty"
)

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt: ">",
		AutoComplete: shellCompleter(),
	})	
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	shell := os.Getenv("SHELL")
	c := exec.Command(shell)
	ptmx, err := pty.Start(c)

	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = ptmx.Close() }()

	go func() {
		_, _ = io.Copy(os.Stdout, ptmx)
	}()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF, readline.ErrInterrupt
			break
		}

		_, err = fmt.Fprintln(ptmx, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	_ = c.Wait()
}

// auto-complete for shell
func shellCompleter() *readline.PrefixCompleter {
	return readline.NewPrefixCompleter(
		readline.PcItem("echo"),
		readline.PcItem("ls",
			readline.PcItem("-la"),
		),
		readline.PcItem("cd",
			readline.PcItem(".."),
			readline.PcItem("/"),
		),
	)
}
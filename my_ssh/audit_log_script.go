package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/creack/pty"
	"golang.org/x/term"
)

const TIMEOUT_PATH = "http://127.0.0.1:8080/api/v1/data/controller/core/aaa/session-config/expire-after-access"

var session_cookie string

func Script(chDone chan bool) error {

	// get shell environment
	// shell := os.Getenv("SHELL")
	// c := exec.Command(shell)
	c := exec.Command("./custom_zsh")

	// start a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	defer func() { _ = ptmx.Close() }()

	// handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH
	defer func() { signal.Stop(ch); close(ch) }()

	// set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }()

	// Read from stdin and send to inputChannel
	inputChannel := make(chan []byte)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				log.Printf("error reading from stdin: %s", err)
				break
			}
			inputChannel <- buf[:n]
		}
	}()

	// Read from ptmx and send to stdout
	go func() {
		_, err := io.Copy(os.Stdout, ptmx)
		if err != nil {
			if err == io.EOF {
				// notify the main loop to exit when ptmx closes
				chDone <- true
				return
			}
			log.Printf("error copying from ptmx to stdout: %s", err)
		}
		chDone <- true
	}()

	// get session cookie
	if session_cookie == "" {
		session_cookie = os.Getenv("FL_SESSION_COOKIE")
		if session_cookie == "" {
			return fmt.Errorf("internal failure: FL_SESSION_COOKIE environment variable is not set")
		}
	}

	// timeout_res, _ := RestGetRequest(session_cookie, TIMEOUT_PATH, "")
	// timeout_min, err := strconv.Atoi(timeout_res)
	if err != nil {
		log.Printf("error transferring from string to int: %s", err)
	}
	// copy stdin to the pty and the pty to stdout.
	for {
		select {
		case input := <-inputChannel:
			_, _ = ptmx.Write(input)
		// use the expiration time of the auth-token to set the timeout
		case <-time.After(5 * time.Second):
			chDone <- true
			return nil
		}
	}
}

func main() {
	flag.StringVar(&session_cookie, "C", "", "Session Cookie Value")
	flag.Parse()
	chDone := make(chan bool)
	go Script(chDone)
	<-chDone
}

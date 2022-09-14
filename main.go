package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/mattn/go-tty"
)

func main() {

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	log.Println("setting up speach-to-text")
	// speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}
	// speech := htgotts.Speech{Folder: "audio", Language: "en"}

	// log.Println("saying hello world")
	// err = speech.Speak("hello, world!")
	// if err != nil {
	// panic(err)
	// }
	// log.Println("said hello world")
	// speech.Speak("listening to tty")

	log.Println("running tty")
	ttyCh := make(chan rune)
	runTTY(tty, ttyCh)

	// log.Println("setting up signal watcher")
	// sigCh := make(chan os.Signal, 1)
	// signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSEGV)

	// speech.Speak("setup complete")

	log.Println("setup complete")
	for {
		select {
		case ttyRune := <-ttyCh:
			// log.Println("Key press => " + string(ttyRune))
			fmt.Println(string(ttyRune))
			// log.Println("setting up speach-to-text")
			// speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}
			// log.Println("speach-to-text setup complete")

			// speech.Speak(string(ttyRune))

			// log.Println("key speech done")
			err = speak(string(ttyRune))
			if err != nil {
				panic(err)
			}

			// case sig := <-sigCh:
			// 	switch sig {
			// 	case os.Interrupt:
			// 		log.Printf("got signal: %s, closing tty", sig)
			// 		tty.Close()
			// 		log.Println("goodbye")
			// 		return
			// 	default:
			// 		log.Printf("got signal: %s, ignoring", sig)
			// 	}
		}
	}
}

func runTTY(tty *tty.TTY, ttyCh chan rune) {
	go func() {
		for {
			r, err := tty.ReadRune()
			if err != nil {
				log.Fatal(err)
			}
			ttyCh <- r
		}
	}()
}

func speak(text string) error {
	out, err := exec.Command("say", text).Output()
	if err != nil {
		return err
	}
	if string(out) != "" {
		log.Printf("output: %s", out)
	}
	return nil
}

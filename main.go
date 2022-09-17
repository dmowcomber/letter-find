package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/dmowcomber/fyne/v2"
	"github.com/dmowcomber/fyne/v2/app"
	"github.com/dmowcomber/fyne/v2/container"
	"github.com/dmowcomber/fyne/v2/widget"
)

func main() {
	var err error
	var minWidth float32 = 500.0
	var minHeight float32 = 500.0

	a := app.New()
	w := a.NewWindow("Hello")
	w.SetPadded(false)
	hello := widget.NewLabel("Letter Find")
	w.SetContent(container.NewVBox(
		hello,
	))
	w.Resize(fyne.NewSize(minWidth, minHeight))

	go func() {
		letter := 'A'
		lastLetter := 'Z'
		hello.SetText("Find the letter")
		fmt.Println("can you find the letter " + string(letter))
		err = speak("can you find the letter " + string(letter))
		if err != nil {
			panic(err)
		}
		lastQuestionTime := time.Now()
		w.Canvas().SetOnTypedRuneEvent(func(runeEvent *fyne.RuneEvent) {
			if runeEvent.When().Before(lastQuestionTime.Add(-1 * time.Second)) {
				hello.SetText("keys pressed during the question are ignored. guess again")
				return // ignore rune keys that were pressed during the question
			}
			eventRune := runeEvent.Rune()

			hello.SetText(fmt.Sprintf("Last letter: %s", strings.ToUpper(string(eventRune))))
			if string(eventRune) == strings.ToUpper(string(letter)) ||
				string(eventRune) == strings.ToLower(string(letter)) {
				err = speak(fmt.Sprintf("correct! that was %s!", string(letter)))
				if err != nil {
					panic(err)
				}
				letter++
				if letter <= lastLetter {
					err = speak("can you find the letter " + string(letter))
					if err != nil {
						panic(err)
					}
					lastQuestionTime = time.Now()
				} else {
					speak("Great job! you got all the letters!")
					os.Exit(0)
				}
			} else {
				err = speak(fmt.Sprintf("no, that was %s. find %s", string(eventRune), string(letter)))
				if err != nil {
					panic(err)
				}
				lastQuestionTime = time.Now()
				return
			}
		})
	}()

	w.ShowAndRun()
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

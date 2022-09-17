package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type runeEvent struct {
	char      rune
	eventTime time.Time
}

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

	// letter := 'A'
	// lastLetter := 'Z'
	lastQuestionTime := time.Now()
	// ignorePrompt := true
	runeEvents := make(chan *runeEvent, 1000)

	// prompt
	go func() {
		for letter := 'A'; letter <= 'Z'; letter++ {
			hello.SetText("Find the letter")
			fmt.Println("can you find the letter " + string(letter))
			err = speak("can you find the letter " + string(letter))
			if err != nil {
				panic(err)
			}
			lastQuestionTime = time.Now()
			for {
				// ignorePrompt = false

				log.Println("fetching input")
				runeEv := <-runeEvents

				eventRune := runeEv.char

				if runeEv.eventTime.Before(lastQuestionTime.Add(-300 * time.Millisecond)) {
					hello.SetText("keys pressed during the question are ignored. guess again")
					continue // ignore rune keys that were pressed during the question
				}

				hello.SetText(fmt.Sprintf("Last letter: %s", strings.ToUpper(string(eventRune))))
				if string(eventRune) == strings.ToUpper(string(letter)) ||
					string(eventRune) == strings.ToLower(string(letter)) {
					err = speak(fmt.Sprintf("correct! that was %s!", string(letter)))
					if err != nil {
						panic(err)
					}
					goto nextLetter
				} else {
					err = speak(fmt.Sprintf("no, that was %s. find %s", string(eventRune), string(letter)))
					if err != nil {
						panic(err)
					}
					lastQuestionTime = time.Now()
					continue
				}
			}
		nextLetter:
		}
		speak("Great job! you got all the letters!")
	}()

	// input
	go func() {
		w.Canvas().SetOnTypedRune(func(char rune) {
			log.Println("writing input to a chan")
			runeEvents <- &runeEvent{
				char:      char,
				eventTime: time.Now(),
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

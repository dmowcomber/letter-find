package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err = s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer func() {
		s.Fini()
	}()

	s.SetStyle(tcell.StyleDefault)

	style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue).Background(tcell.ColorWhite)

	// log.Println("setup complete")

	w, h := s.Size()
	s.Clear()
	emitStr(s, w/2-7, h/2, style, "setup complete")
	s.Show()

	// TODO: randomize letters?
	for letter := 'a'; letter <= 'z'; letter++ {
		err = speak("can you find the letter " + string(letter))
		if err != nil {
			panic(err)
		}
		lastQuestionTime := time.Now()

		for {
			// TODO: why did the example do this in a goroutine?
			// https://github.com/gdamore/tcell/blob/22d72263215d7b0298d6d4881053b042192117a7/_demos/cursors.go#L53
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
					return
				case tcell.KeyRune:
					if ev.When().Before(lastQuestionTime) {
						continue // ignore rune keys that were pressed during the question
					}

					eventRune := ev.Rune()
					// s.SetCell(2, 2, tcell.StyleDefault, '6')
					w, h := s.Size()
					s.Clear()
					// emitStr(s, w/2-7, h/2, style, "Hello, World!")
					emitStr(s, w/2-7, h/2, style, "Key press => "+string(eventRune))
					s.Show()
					// log.Println("Key press => " + string(eventRune))
					if string(eventRune) == string(letter) {
						err = speak(fmt.Sprintf("correct! that was %s!", string(letter)))
						if err != nil {
							panic(err)
						}
						goto NextLetter
					} else {
						err = speak(fmt.Sprintf("no, that was %s. find %s", string(eventRune), string(letter)))
						if err != nil {
							panic(err)
						}
						lastQuestionTime = time.Now()
					}
				}
			}
		}
	NextLetter:
	}
	speak("Great job! you got all the letters!")
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
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

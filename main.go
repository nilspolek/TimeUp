package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mbndr/figlet4go"
	"github.com/nilspolek/TimeUp/box"
)

var color string
var isLong bool

func main() {
	flag.StringVar(&color, "color", "Blue", "gookit/colors (Blue, Black, Red, Green, ...)")
	flag.BoolVar(&isLong, "l", false, "display with milliseconds")
	flag.Parse()
	Box := box.New(box.Config{Px: 2, Py: 2, Type: "Round", Color: color, TitlePos: "Top"})
	tm := time.Now()
	ascii := figlet4go.NewAsciiRender()

	for {
		since := time.Since(tm)
		hours := int(since.Hours()) % 24
		minutes := int(since.Minutes()) % 60
		secounds := int(since.Seconds()) % 60
		mills := int(since.Milliseconds()) % 1000
		renderStr := ""
		if isLong {
			renderStr, _ = ascii.Render(fmt.Sprintf("%02d:%02d:%02d:%03d", hours, minutes, secounds, mills))
		} else {
			renderStr, _ = ascii.Render(fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secounds))
		}
		Box.PrintAlingned("TimeUp", renderStr)
		time.Sleep(time.Millisecond * 10)
	}
}

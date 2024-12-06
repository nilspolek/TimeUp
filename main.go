package main

import (
	"fmt"
	"time"

	"github.com/mbndr/figlet4go"
	"github.com/nilspolek/TimeUp/box"
)

func main() {
	Box := box.New(box.Config{Px: 1, Py: 1, Type: "Single", Color: "Cyan", TitlePos: "Top"})
	tm := time.Now()
	ascii := figlet4go.NewAsciiRender()

	for {
		since := time.Since(tm)
		renderStr, _ := ascii.Render(fmt.Sprintf("%02d:%02d:%02d", int(since.Hours())%24, int(since.Minutes())%60, int(since.Seconds())%60))
		Box.PrintAlingned("TimeUp", renderStr)
		time.Sleep(time.Millisecond * 10)
	}
}

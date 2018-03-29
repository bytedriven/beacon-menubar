package main

import (
    "fmt"
    "io/ioutil"

    "github.com/getlantern/systray"
   // "github.com/mikepb/go-serial"
)

var (
    title string
)

func main() {
    systray.Run(onReady, onExit)
}

func onReady() {
    title = "Available"
    systray.SetIcon(getIcon("assets/circle.ico"))
    systray.SetTitle(title)
    systray.SetTooltip("this is a tolltip")

    available := systray.AddMenuItem("Available", "Set yourself as ready to mingle")
    busy := systray.AddMenuItem("Busy", "Don't bother me, I'm busy")
    away := systray.AddMenuItem("Away", "Up up and away")
    systray.AddSeparator()
    // Going to list all serial devices here
    quit := systray.AddMenuItem("Quit", "Let's get out of here")

    go func() {
        for {
            systray.SetTitle(title)
        }
    }()

    go func() {
        for {
            select {
                case <-available.ClickedCh:
                    title = "Available"
                case <-busy.ClickedCh:
                    title = "Busy"
                case <-away.ClickedCh:
                    title = "Away"
                case <-quit.ClickedCh:
                    systray.Quit()
                    return
            }
        }
    }()
}

func onExit() {
    //cleanup
}

func getIcon(s string) []byte {
    b, err := ioutil.ReadFile(s)
    if err != nil {
        fmt.Print(err)
    }
    return b
}

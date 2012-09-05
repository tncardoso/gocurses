package main

import (
    "fmt"
    "github.com/tncardoso/gocurses"
)

const (
    windowHeight = 20
    windowWidth  = 40
)

func main() {
    fmt.Println("starting")
    gocurses.Initscr()
    defer gocurses.End()
    gocurses.Cbreak()
    gocurses.Noecho()
    gocurses.Stdscr.Keypad(true)

    y, x := gocurses.Getmaxyx()
    gocurses.Addstr("Press any key to exit")
    gocurses.Refresh()

    window := gocurses.NewWindow(windowHeight, windowWidth, (y-windowHeight)/2, (x-windowWidth)/2)
    window.Box(0, 0)
    window.Mvaddstr(0, 1, "Sample")
    window.Refresh()

    gocurses.Getch()
}

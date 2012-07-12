package main

import "github.com/tncardoso/gocurses"

const (
    windowHeight = 20
    windowWidth  = 40
)

func main() {
    curses.Initscr()
    defer curses.End()
    curses.Cbreak()
    curses.Noecho()
    curses.Stdscr.Keypad(true)

    y, x := curses.Getmaxyx()
    curses.Addstr("Press any key to exit")
    curses.Refresh()

    window := curses.NewWindow(windowHeight, windowWidth, (y-windowHeight)/2, (x-windowWidth)/2)
    window.Box(0, 0)
    window.Mvaddstr(0, 1, "Sample")
    window.Refresh()

    curses.Getch()
}

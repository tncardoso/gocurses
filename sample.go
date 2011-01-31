package main

import "github.com/thiagoncc/curses"

func main() {
    curses.Initscr()
    curses.Cbreak()
    curses.Noecho()
    curses.Stdscr.Keypad(true)
    curses.Attron(curses.A_BOLD)
    curses.Printw("Having big lols: ", "LOLG")
    wind := curses.NewWindow(10,40,10,10)
    wind.Box(0,0)
    x, y := curses.Getmaxyx()
    curses.Refresh()
    if curses.Stdscr.Getch() == curses.KEY_DOWN {
        wind.Mvprintw(1,1,"X; ",x, " Y: ",y)
        wind.Refresh()
        curses.Stdscr.Getch()
    }
    curses.End()
}

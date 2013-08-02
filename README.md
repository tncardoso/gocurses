curses.go
=============

GO binding for ncurses.

Sample
-------
    package main

    import "github.com/tncardoso/gocurses"

    func main() {
        gocurses.Initscr()
        defer gocurses.End()
        gocurses.Cbreak()
        gocurses.Noecho()
        gocurses.Stdscr.Keypad(true)
        
        gocurses.Attron(gocurses.A_BOLD)
        gocurses.Addstr("Hello World!")
        gocurses.Refresh()

        wind := gocurses.NewWindow(10,40,10,10)
        wind.Box(0,0)
        wind.Refresh()
        
        gocurses.Stdscr.Getch()
    }

Requirements
-------

* [libncurses](http://ftp.gnu.org/pub/gnu/ncurses/) -- ncurses library

Installation
-------

goinstall is now supporting cgo packages, therefore installing gocurses
should be as easy as:

    $ go get github.com/tncardoso/gocurses

If you have problems with go get, you can update your go release or
clone the repository:

    $ git clone git://github.com/tncardoso/gocurses.git
    $ cd gocurses
    $ go install
curses.go
=============

GO binding for ncurses.

Sample
-------
    func main() {
        curses.Initscr()
        defer curses.End()
        curses.Cbreak()
        curses.Noecho()
        curses.Stdscr.Keypad(true)
        
        curses.Attron(curses.A_BOLD)
        curses.Printw("Hello World!")
        curses.Refresh()

        wind := curses.NewWindow(10,40,10,10)
        wind.Box(0,0)
        wind.Refresh()
        
        curses.Stdscr.Getch()
    }

Requirements
-------

* [libncurses](http://ftp.gnu.org/pub/gnu/ncurses/) -- ncurses library

Instalation
-------

goinstall is now supporting cgo packages, therefore installing curses.go
should be as easy as:

    $ goinstall github.com/thiagoncc/curses.go

If you have problems with goinstall, you can update your go release or
clone the repository:

    $ git clone git://github.com/thiagoncc/curses.go.git
    $ cd curses.go
    $ gomake install

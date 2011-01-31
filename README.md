curses.go
=============

GO binding for ncurses.

Sample
-------
    func main() {
        curses.Initscr()
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
        curses.End()
    }

Requirements
-------

* [libncurses](http://ftp.gnu.org/pub/gnu/ncurses/) -- ncurses library

Instalation
-------

    $ goinstall github.com/thiagoncc/curses.go

package gocurses

// #cgo LDFLAGS: -lncurses
// #include <stdlib.h>
// #include <ncurses.h>
// int wrapper_getmaxx(WINDOW* win) { return getmaxx(win); }
// int wrapper_getmaxy(WINDOW* win) { return getmaxy(win); }
// void wrapper_wclrtoeol(WINDOW* win) { wclrtoeol(win); }
// void wrapper_wattrset(WINDOW* win, int attr) { wattrset(win, attr); }
//
import "C"
import "unsafe"
import "fmt"

// Curses window type.
type Window struct {
    cwin *C.WINDOW
}

// Standard window.
var Stdscr *Window = &Window{cwin: C.stdscr}

// Initializes curses.
// This function should be called before using the package.
func Initscr() *Window {
    Stdscr.cwin = C.initscr()
    return Stdscr
}

// Raw input. No buffering.
// CTRL+Z and CTRL+C passed to the application.
func Raw() {
    C.raw()
}

// No buffering.
func Cbreak() {
    C.cbreak()
}

// Enable character echoing while reading.
func Echo() {
    C.echo()
}

// Disables character echoing while reading.
func Noecho() {
    C.noecho()
}

func Curs_set(i int) {
  C.curs_set(C.int(i))
}

// Enable reading of function keys.
func (window *Window) Keypad(on bool) {
    C.keypad(window.cwin, C.bool(on))
}

// Get char from the standard in.
func (window *Window) Getch() int {
    return int(C.wgetch(window.cwin))
}

// Get char from the standard in.
func Getch() int {
    return int(C.getch())
}

// Enable attribute
func Attron(attr int) {
    C.attron(C.int(attr))
}

// Disable attribute
func Attroff(attr int) {
    C.attroff(C.int(attr))
}

// Set attribute
func Attrset(attr int) {
	C.attrset(C.int(attr))
}

func (window *Window) Attron(attr int) {
	C.wattron(window.cwin, C.int(attr))
}

func (window *Window) Attroff(attr int) {
	C.wattroff(window.cwin, C.int(attr))
}

func (window *Window) Attrset(attr int) {
	C.wrapper_wattrset(window.cwin, C.int(attr))
}

// Refresh screen.
func Refresh() {
    C.refresh()
}

// Refresh given window.
func (window *Window) Refresh() {
    C.wrefresh(window.cwin)
}

// Finalizes curses.
func End() {
    C.endwin()
}

// Create new window.
func NewWindow(height, width, starty, startx int) *Window {
    w := new(Window)
    w.cwin = C.newwin(C.int(height), C.int(width),
        C.int(starty), C.int(startx))
    return w
}

// Set box lines.
func (window *Window) Box(v, h int) {
    C.box(window.cwin, C.chtype(v), C.chtype(h))
}

// Set border characters.
// 1. ls: character to be used for the left side of the window 
// 2. rs: character to be used for the right side of the window 
// 3. ts: character to be used for the top side of the window 
// 4. bs: character to be used for the bottom side of the window 
// 5. tl: character to be used for the top left corner of the window 
// 6. tr: character to be used for the top right corner of the window 
// 7. bl: character to be used for the bottom left corner of the window 
// 8. br: character to be used for the bottom right corner of the window
func (window *Window) Border(ls, rs, ts, bs, tl, tr, bl, br int) {
    C.wborder(window.cwin, C.chtype(ls), C.chtype(rs), C.chtype(ts), C.chtype(bs), C.chtype(tl), C.chtype(tr), C.chtype(bl), C.chtype(br))
}

// Delete current window.
func (window *Window) Del() {
    C.delwin(window.cwin)
}

// Get windows sizes.
func (window *Window) Getmaxyx() (row, col int) {
    row = int(C.wrapper_getmaxy(window.cwin))
    col = int(C.wrapper_getmaxx(window.cwin))
    return row, col
}

func (window *Window) Setscrreg(top, bot int) {
    C.wsetscrreg(window.cwin, C.int(top), C.int(bot))
}

func Addstr(str ...interface{}) {
    res := (*C.char)(C.CString(fmt.Sprint(str...)))
    defer C.free(unsafe.Pointer(res))
    C.addstr(res)
}

func Mvaddstr(y, x int, str ...interface{}) {
    res := (*C.char)(C.CString(fmt.Sprint(str...)))
    defer C.free(unsafe.Pointer(res))
    C.mvaddstr(C.int(y), C.int(x), res)
}

func Addch(ch int) {
	C.addch(C.chtype(ch))
}

func Mvaddch(y, x int, ch int) {
	C.mvaddch(C.int(y), C.int(x), C.chtype(ch))
}

func (window *Window) Addstr(str ...interface{}) {
    res := (*C.char)(C.CString(fmt.Sprint(str...)))
    defer C.free(unsafe.Pointer(res))
    C.waddstr(window.cwin, res)
}

func (window *Window) Mvaddstr(y, x int, str ...interface{}) {
    res := (*C.char)(C.CString(fmt.Sprint(str...)))
    defer C.free(unsafe.Pointer(res))
    C.mvwaddstr(window.cwin, C.int(y), C.int(x), res)
}

func (window *Window) Addch(ch int) {
	C.waddch(window.cwin, C.chtype(ch))
}

func (window *Window) Mvaddch(y, x int, ch int) {
	C.mvwaddch(window.cwin, C.int(y), C.int(x), C.chtype(ch))
}

// Hardware insert/delete feature.
func (window *Window) Idlok(bf bool) {
    C.idlok(window.cwin, C.bool(bf))
}

// Enable window scrolling.
func (window *Window) Scrollok(bf bool) {
    C.scrollok(window.cwin, C.bool(bf))
}

// Scroll given window.
func (window *Window) Scroll() {
    C.scroll(window.cwin)
}

// Get terminal size.
func Getmaxyx() (row, col int) {
    row = int(C.LINES)
    col = int(C.COLS)
    return row, col
}

// Erases content from cursor to end of line inclusive.
func (window *Window) Clrtoeol() {
    C.wrapper_wclrtoeol(window.cwin)
}

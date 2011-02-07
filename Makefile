include $(GOROOT)/src/Make.inc

TARG=github.com/thiagoncc/curses
CGOFILES=curses.go curses_defs.go
CGO_LDFLAGS=-lncurses

include $(GOROOT)/src/Make.pkg

format:
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w curses.go
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w curses_defs.go


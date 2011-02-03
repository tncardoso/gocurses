include $(GOROOT)/src/Make.inc

TARG=github.com/thiagoncc/curses
CGOFILES=curses.go curses_defs.go
CGO_LDFLAGS=-lncurses

include $(GOROOT)/src/Make.pkg

sample: *.go
	$(GC) sample.go
	$(LD) -o sample sample.$(O)


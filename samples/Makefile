include $(GOROOT)/src/Make.inc

simple: *.go
	$(GC) simple.go
	$(LD) -o simple simple.$(O)

format:
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w simple.go

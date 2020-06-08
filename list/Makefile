.PHONY: clean all

all: function.zip

list: main.go
	go build

function.zip: list
	zip function.zip list

clean:
	rm -f list function.zip

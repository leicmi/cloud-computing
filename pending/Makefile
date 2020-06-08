.PHONY: clean all

all: function.zip

pending: main.go
	go build

function.zip: pending
	zip function.zip pending

clean:
	rm -f pending function.zip

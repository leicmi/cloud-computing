.PHONY: clean all

all: function.zip

upload: main.go
	go build

function.zip: upload
	zip function.zip upload

clean:
	rm -f upload function.zip

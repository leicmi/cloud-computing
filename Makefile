.PHONY: clean lamq

all: lamq
	$(MAKE) -C upload all

lamq:
	go build -o lamq

clean:
	rm -rf lamq
	$(MAKE) -C upload clean

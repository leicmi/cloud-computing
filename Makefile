.PHONY: clean lamq

all: lamq
	$(MAKE) -C upload all
	$(MAKE) -C pending all
	$(MAKE) -C list all

lamq:
	go build -o lamq

clean:
	rm -rf lamq
	$(MAKE) -C upload clean
	$(MAKE) -C pending clean
	$(MAKE) -C list clean

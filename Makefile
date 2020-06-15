.PHONY: clean lamq

all: lamq
	$(MAKE) -C "AWS Lambdas" build

lamq:
	go build -o lamq

clean:
	rm -rf lamq
	$(MAKE) -C "AWS Lambdas" clean

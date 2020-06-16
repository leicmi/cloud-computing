.PHONY: build deploy clean

build:
	sam build

deploy:
	sam deploy -g

clean:
	rm -rf .aws-sam/build
	rmdir .aws-sam

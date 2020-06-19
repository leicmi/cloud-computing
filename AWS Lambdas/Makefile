.PHONY: build deploy fastdeploy clean

build:
	sam build

deploy:
	sam deploy -g

fastdeploy:
	sam deploy

clean:
	rm -rf .aws-sam/build
	rmdir .aws-sam

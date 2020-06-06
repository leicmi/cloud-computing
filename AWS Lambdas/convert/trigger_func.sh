#!/bin/bash

user=$1
sourcebin=$2

if [ -z "$user" ] || [ -z "$sourcebin" ]; then 
	echo usage: $0 user sourcebin
        exit
fi

#check if function is already uploaded, upload if necessary
aws lambda get-function --function-name $CreateThumbnail > /dev/null 2>&1
  if [ 0 -eq $? ]; then
    echo "Lambda '$1' exists, now writing permissions"
  else
    echo "Lambda '$1' does not exist, downloading files to create"
	instl=(npm --version)
	if [ $instl != 6.14.4 ]; then
	echo "npm is required and not installed, give permission to download"
	curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
	sudo apt install nodejs
	fi

	npm install sharp
	zip -r function.zip .

	echo "creating function on AWS lambda"
	aws lambda create-function --function-name CreateThumbnail \
	--zip-file fileb://function.zip --handler index.handler --runtime nodejs12.x \
	--timeout 10 --memory-size 1024 \
	--role arn:aws:iam::$user:role/lambda-s3-role
	#optional timeout change
	#aws lambda update-function-configuration --function-name CreateThumbnail --timeout 30

  fi


aws lambda add-permission --function-name CreateThumbnail --principal s3.amazonaws.com \
--statement-id s3invoke --action "lambda:InvokeFunction" \
--source-arn arn:aws:s3:::$sourcebin \
--source-account $user



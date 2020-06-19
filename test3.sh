#!/bin/bash

## performance

BUCKET="leicmi-cloud-computing-lamqbucket-1ulbttrj6irl3"

#aws s3 cp ./image.png s3://$BUCKET/image.png
#sleep 10

#while true
for i in $(seq 1 100)
do
    for j in $(seq 1 100)
    do
        curl --silent --output nul --show-error --fail "https://d46t21t0m3.execute-api.eu-central-1.amazonaws.com/Prod/trigger" &
    done
    sleep 1
done

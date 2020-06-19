#!/bin/bash

## performance

BUCKET="leicmi-cloud-computing-lamqbucket-1ulbttrj6irl3"

aws s3 cp ./image.png s3://$BUCKET/image.png
sleep 10

#while true
for i in $(seq 1 100)
do
    for j in $(seq 1 100)
    do
        aws s3 cp s3://$BUCKET/image.png s3://$BUCKET/image1-$i.png &
    done
    sleep 3
done

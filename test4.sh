#!/bin/bash

## elasticity
## TODO: redeploy before!

BUCKET="leicmi-cloud-computing-lamqbucket-1ulbttrj6irl3"

aws s3 cp ./image.png s3://$BUCKET/image.png
sleep 10

#while true
for i in $(seq 1 1000)
do
    aws s3 cp s3://$BUCKET/image.png s3://$BUCKET/image4-$i.png &
done

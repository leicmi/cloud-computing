#!/bin/bash

#while true
for i in $(seq 1 200)
do
    aws s3 cp s3://leicmi-com-github-leicmi-cloud-computi-lamqbucket-i0ivzbcgzc7f/image.png s3://leicmi-com-github-leicmi-cloud-computi-lamqbucket-i0ivzbcgzc7f/image1-$i.png
    sleep 2
done
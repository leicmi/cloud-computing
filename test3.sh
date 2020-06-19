#!/bin/bash

## performance

BUCKET="leicmi-cloud-computing-lamqbucket-1ulbttrj6irl3"

#aws s3 cp ./image.png s3://$BUCKET/image.png
sleep 10

#while true
for i in $(seq 1 10)
do
    for j in $(seq 1 1)
    do
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-10.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-11.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-12.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-13.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-14.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-15.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-16.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-17.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-18.png s3://$BUCKET/image3-100-$i-$j.png &
        aws s3 cp --quiet --only-show-errors --no-guess-mime-type s3://$BUCKET/image2-3008-19.png s3://$BUCKET/image3-100-$i-$j.png &
    done
    sleep 1
done

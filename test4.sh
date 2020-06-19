#!/bin/bash

## elasticity
## TODO: redeploy before!

for i in $(seq 1 100)
do
    curl --silent --output nul --show-error --fail "https://d46t21t0m3.execute-api.eu-central-1.amazonaws.com/Prod/trigger" &
done

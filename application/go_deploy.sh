#!/bin/bash
#



echo "deploy go app"

echo "building"

sudo go build main.go


nohup ./main > go_deploy.log 2>&1 &


echo "depoly successful"

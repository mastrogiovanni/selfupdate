#!/bin/bash

echo Building software

go build -o old1 old/main.go

go build -o new2 new/main.go

bsdiff old1 new2 new1

echo Run old software:
echo it will download and update itself
echo the payload downloaded is only the bsdiff 
echo between current version and new version

./old1

echo Run the same software that is correctly updated

./old1


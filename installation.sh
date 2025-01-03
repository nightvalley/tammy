#!/bin/bash

go build -o tammy cmd/main.go

if [ -f "/usr/bin/tammy" ]; then
  sudo rm /usr/bin/tammy
  sudo mv tammy /usr/bin
else
  sudo mv tammy /usr/bin
fi

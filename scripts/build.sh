#!/bin/bash
go get -v -t ./...
go build -o ./build/unicore_network -v github.com/praekelt/unicore_network

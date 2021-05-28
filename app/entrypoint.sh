#!/bin/sh

set -eu

go build -o app
./app

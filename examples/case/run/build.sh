#!/bin/bash

set -e

go build -o bin/case ./pkg
./bin/case --manifest case -location plugin/case.vim



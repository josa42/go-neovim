#!/bin/bash

set -e

uuid="b33d0c0da04fae2980d7f160c26bd50b"

go build -ldflags "-X main.uuid=${uuid}" -o bin/case ./pkg
./bin/case --manifest case -location plugin/case.vim



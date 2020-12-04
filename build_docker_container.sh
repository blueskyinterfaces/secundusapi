#!/bin/bash

(cd ~/go/src/github.com/secundusteam/secundus/assets/app && yarn && yarn lint --fix && yarn build)

(cd ~/go/src/github.com/secundusteam/secundus && docker build -t secundus .)
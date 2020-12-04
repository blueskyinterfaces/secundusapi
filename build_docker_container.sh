#!/bin/bash

(cd ~/go/src/github.com/blueskyinterfaces/secundusapiapiapi/assets/app && yarn && yarn lint --fix && yarn build)

(cd ~/go/src/github.com/blueskyinterfaces/secundusapiapiapi && docker build -t secundusapi .)
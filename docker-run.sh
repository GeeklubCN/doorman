#!/bin/bash

docker run -v $PWD/docker-custom-conf.yaml:/go/src/github.com/geeklubcn/doorman/conf/config.yaml:ro wangyuheng/doorman:v1
#!/bin/bash
docker kill code-runner
docker rm code-runner
docker run --net helloclass_hello-class-network -ditp 8079:8079 --name code-runner code-runner
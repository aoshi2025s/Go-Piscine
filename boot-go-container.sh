# !/bin/bash

docker run -it -v $(pwd)/:/usr/src/app -w /usr/src/app/ golang:latest bash

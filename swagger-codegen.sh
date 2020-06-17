#!/bin/sh

alias swagger="docker run --network host --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"

#curl http://localhost:5002/swagger/v1/swagger.json > api-swagger.json
#docker run -it --rm -v $(pwd):/data swaggerapi/swagger-codegen-cli generate -i data/api-swagger.json -l go -o /data/swagger && \
swagger generate client -f http://localhost:5002/swagger/v1/swagger.json -A deply-f-api --skip-validation
sudo chown -R $USER:$USER client
sudo chown -R $USER:$USER models
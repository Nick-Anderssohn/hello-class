# intermediate container used to build the dart/scss code
FROM nickanderssohn/dart-sass:1.0 as client-builder

# Copy project over
RUN mkdir /client
WORKDIR /client
COPY ./client .

# build dart
RUN pub get
RUN pub build
RUN cp -R build/web/packages/codemirror build/web/play/packages/
RUN cp -R build/web/packages/codemirror build/web/question/packages/

# The final container
FROM golang:1.9.0-alpine
EXPOSE 8080
EXPOSE 8079
WORKDIR /go/src
RUN apk add --update curl
RUN rm -rf /var/cache/apk/*
RUN mkdir hello-compsci
RUN mkdir hello-compsci/frontend-server
WORKDIR /go/src/hello-compsci/frontend-server
COPY . .
RUN rm -rf client && mkdir client && mkdir client/build
# copy the version that was compiled in the intermediate container
COPY --from=client-builder /client/build ./client/build
# RUN apk update && apk upgrade && apk add git
# RUN go get "golang.org/x/crypto/acme/autocert"
RUN go install
WORKDIR /go/src/hello-compsci/frontend-server/client
WORKDIR /go/src/hello-compsci/frontend-server
RUN rm Dockerfile
# commented out healthcheck because it appears to have issues with https
# HEALTHCHECK CMD curl --fail https://localhost:8080/ || exit 1
#COPY selfgen .
CMD frontend-server
##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /go_app
#RUN apt update -y
#RUN apt install -y tree 

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./go_app ./

RUN go build -o ./go_app

#RUN ls
#RUN tree

##
## Deploy
##
FROM ubuntu
RUN useradd -ms /bin/bash golaner
USER golaner
WORKDIR /
#RUN apt update -y
#RUN apt install -y tree 

COPY --from=build --chown=golaner:golaner /go_app /go_app

EXPOSE 8080

WORKDIR /go_app

RUN chmod +x go_app

ENTRYPOINT ["./go_app"]
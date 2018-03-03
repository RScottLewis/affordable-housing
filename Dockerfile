FROM golang:1.9-stretch
MAINTAINER r.scott.lewis@gmail.com

ARG  CODE_VERSION=latest
#FROM base:${CODE_VERSION}
#CMD  go run hellosrv.go
CMD git clone https://github.com/RScottLewis/affordable-housing;go run affordable-housing/affordable-housing.go


RUN mkdir -p /go/src/app
#CMD mkdir -p /go/src/app

#this will ideally be built by the ONBUILD below :)
#CMD["go-wrapper"]

ONBUILD ADD . /go/src/app
WORKDIR /go/src/app

#ONBUILD COPY /Volumes/ITM-STICK2/Briefcase4NOV17/src/go/hellosrv/* /go/src/app
ONBUILD COPY . /go/src/app
#RUN cp /Volumes/ITM-STICK2/Briefcase4NOV17/src/go/hellosrv/* /go/src/app
#ONBUILD RUN /usr/local/bin/python-build --dir /app/src
ONBUILD RUN go-wrapper download
ONBUILD RUN go-wrapper install

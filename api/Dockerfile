# API
FROM golang:1.12.5-stretch AS development

# ENV GO111MODULE=on
RUN go get github.com/pilu/fresh
WORKDIR /source/api

EXPOSE 80
ENTRYPOINT ["fresh"]

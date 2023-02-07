FROM golang:1.20rc1

ENV GOPATH /go/src

RUN mkdir /go/src/backend
WORKDIR /go/src/backend

COPY go.mod /go/src/backend/go.mod
COPY go.sum /go/src/backend/go.sum

RUN go mod download

# RUN go install \ 
#     github.com/cweill/gotests/gotests \
#     github.com/fatih/gomodifytags \
#     github.com/josharian/impl \
#     github.com/haya14busa/goplay/cmd/goplay \
#     github.com/go-delve/delve/cmd/dlv \
#     honnef.co/go/tools/cmd/staticcheck \
#     golang.org/x/tools/gopls
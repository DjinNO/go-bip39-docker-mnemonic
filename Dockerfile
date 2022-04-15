FROM golang:alpine3.15 as build
RUN apk update && apk add git
COPY main.go .
ENV GO111MODULE=off 
RUN go get github.com/tyler-smith/go-bip39 && go build -o /mnemonic

FROM golang:alpine3.15
COPY --from=build /mnemonic /mnemonic
ENTRYPOINT ["/mnemonic"]
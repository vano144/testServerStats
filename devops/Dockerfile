FROM golang:1.12 as build-env

RUN mkdir /src
WORKDIR /src
COPY go.mod go.sum ./

RUN go mod download
COPY . .

# Build the binary
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install  ./... # build all where we find main function and locate it in /go/bin
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -i -o /go/bin/testServerStats-$GOOS-$GOARCH ./cmd/statistics/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test  ./...

FROM alpine:latest AS runtime
RUN apk --no-cache add ca-certificates
COPY --from=build-env /go/bin/testServerStats-$GOOS-$GOARCH /go/bin/testServerStats-$GOOS-$GOARCH
ENTRYPOINT ./go/bin/testServerStats-$GOOS-$GOARCH
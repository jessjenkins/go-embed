FROM golang:1.16 as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -json
RUN go list -f '{{.Path}}/...' -m all | tail -n +2 | CGO_ENABLED=0 GOOS=linux xargs -n1 go build -v -installsuffix cgo -i
COPY . .
RUN find . -type f && CGO_ENABLED=0 GOOS=linux go build -a -v -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build app/app .
RUN mkdir -p files/sub \
  && echo "Woof" > files/woof.txt \
  && echo "Meow" > files/sub/meow.txt
CMD ["./app"]

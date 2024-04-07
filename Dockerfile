FROM golang:alpine AS build

WORKDIR /go/src/filehub

COPY . .

RUN go build -o /go/bin/filehub cmd/main.go

FROM scratch
COPY --from=build /go/bin/filehub /go/bin/filehub
EXPOSE 8080
ENTRYPOINT ["/go/bin/filehub"]

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go ./

RUN go build -o ./gosum

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/gosum /gosum

USER nonroot:nonroot

ENTRYPOINT ["./gosum"]
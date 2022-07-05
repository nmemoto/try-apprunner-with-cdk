# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /app

##
## Deploy
##
FROM gcr.io/distroless/base-debian11
ENV PORT=8080
WORKDIR /
COPY --from=build /app /
EXPOSE ${PORT}
USER nonroot:nonroot
ENTRYPOINT ["/app"]
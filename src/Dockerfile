FROM golang:1.11-alpine AS build
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /src
ADD . /src
RUN go mod download
RUN go build -o app

FROM alpine
WORKDIR /app
COPY --from=build /src/app /app/
ENTRYPOINT ./app
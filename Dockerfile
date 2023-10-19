FROM golang:1.21.3-alpine AS build

WORKDIR /go/src/app
COPY . .
ENV GOOS=linux
RUN go build -v -o /go/bin/app .

FROM alpine:3.18
WORKDIR /app
COPY --from=build /go/bin/app .
CMD ["./app"]

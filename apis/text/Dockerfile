FROM golang:1.19 as builder

WORKDIR /app

COPY . ./

RUN go build -o ./build/program .

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/build/program ./program

ENTRYPOINT [ "./program" ]
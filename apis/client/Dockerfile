FROM golang:1.19 as builder

WORKDIR /app

COPY . ./

RUN go build -o ./build/program .

FROM ubuntu:22.10

WORKDIR /app

COPY --from=builder /app/build/program ./program
COPY --from=builder /app/deployment.template.json ./deployment.template.json

ENTRYPOINT [ "./program" ]
FROM golang:1.19

WORKDIR /app

COPY . ./

RUN go build -o ./program .
RUN if [ ! -d download ]; then \
    mkdir download; \
    fi

ENTRYPOINT [ "./program" ]
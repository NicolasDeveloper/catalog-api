FROM golang:latest

WORKDIR /app

ARG LOG_DIR=/app/logs

RUN mkdir -p ${LOG_DIR}

ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main /main

EXPOSE 3201

VOLUME [${LOG_DIR}]

CMD ["./main"]
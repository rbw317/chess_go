# syntax=docker/dockerfile:1

FROM golang:1.21
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
WORKDIR /app/chess_engine
COPY chess_engine ./
WORKDIR /app/chess_web_page
COPY chess_web_page ./
WORKDIR /app/chess_web_service
COPY chess_web_service ./
WORKDIR /app
RUN go get chess_go/chess_engine
RUN go get chess_go/chess_web_service
RUN CGO_ENABLED=0 GOOS=linux go build -o /chess_go
EXPOSE 80
CMD ["/chess_go"]
FROM golang:1.17-alpine
WORKDIR /app
COPY . .
COPY go.mod go.sum /app/
COPY *.go ./
RUN go build -o /main
EXPOSE 8080
CMD ["/main"]

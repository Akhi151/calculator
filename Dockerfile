FROM golang:latest
RUN mkdir/sub
WORKDIR /sub
COPY . /sub
RUN go mod download
RUN cd/sub/calculator/development
EXPOSE 8000
CMD ["/main.go"]
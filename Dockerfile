FROM golang:latest
RUN mkdir/sub
COPY . /sub
WORKDIR /sub
RUN cd/sub/calculator/development
EXPOSE 8000
CMD ["/main.go"]
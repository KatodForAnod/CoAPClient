FROM golang:latest
RUN apt-get update
RUN apt-get install git -y
WORKDIR /app
RUN git clone https://github.com/KatodForAnod/CoAPServer.git && \
    cd CoAPServer && \
    go mod tidy && \
    go build gocoaplib/iotemulatelite.go
CMD ["/app/CoAPServer/iotemulatelite"]
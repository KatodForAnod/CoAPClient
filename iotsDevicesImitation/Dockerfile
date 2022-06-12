FROM golang:latest
RUN apt-get update
RUN apt-get install git -y
WORKDIR /app
RUN git clone https://github.com/KatodForAnod/CoAPIoTDevice.git --branch v1.0.0 && \
    cd CoAPIoTDevice && \
    go mod tidy && \
    go build gocoaplib/iotemulatelite.go
CMD ["sh", "-c", "/app/CoAPIoTDevice/iotemulatelite -port ${port}"]
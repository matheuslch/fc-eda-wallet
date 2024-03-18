FROM golang:latest

WORKDIR /app/

RUN apt-get update && apt-get install -y curl lsb-release software-properties-common
RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey| apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/debian/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list 
RUN apt-get update && apt-get install -y librdkafka-dev gcc binutils xfsprogs build-essential migrate

ENV GOOS linux
ENV GOARCH amd64

COPY /cmd/walletcore/main /cmd/walletcore/main
RUN chmod +x /cmd/walletcore/main

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]

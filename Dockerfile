FROM golang:1.12.5 AS builder
RUN go version

COPY . /kalista
WORKDIR /kalista
RUN set -x && \
  go get -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o kalista-app .

FROM scratch
WORKDIR /root/

COPY --from=builder /kalista/kalista-app .
# COPY ./config/config.yml /root/config/config.yml

EXPOSE 8080
ENTRYPOINT ["./kalista-app"]
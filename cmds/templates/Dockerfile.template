FROM golang:1.12.7 AS builder

RUN go version
COPY . /{{.ModuleName}}
WORKDIR /{{.ModuleName}}
RUN set -x && \
  go get -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o {{.ModuleName}}-app .

FROM scratch
WORKDIR /root/

COPY --from=builder /{{.ModuleName}}/{{.ModuleName}}-app .
# COPY ./config/config.yml /root/config/config.yml

EXPOSE 8080
ENTRYPOINT ["./{{.ModuleName}}-app"]
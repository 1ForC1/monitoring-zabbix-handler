FROM golang:1.19-alpine

WORKDIR /zabbix

COPY . .

RUN go build -o zabbix-data-handler main.go

RUN chmod +x zabbix-data-handler

CMD ["./zabbix-data-handler"]
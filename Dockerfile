FROM scratch

ENV SERVICE_PORT 8000

EXPOSE $SERVICE_PORT

COPY go-kubernetes /

CMD ["/go-kubernetes"]
FROM golang
ADD . /whereami/
ADD main.go /whereami/
WORKDIR /whereami/
RUN CGO_ENABLED=0 go build /whereami/main.go
EXPOSE 8080
ENTRYPOINT ["./main"]
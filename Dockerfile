FROM shivam010/golang
WORKDIR /app
COPY /cmd/bot/main.go /app/main.go
ENTRYPOINT [ "make" ]


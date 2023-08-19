FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN go build -o /docker-pi-api
EXPOSE 8080
CMD ["/docker-pi-api"]
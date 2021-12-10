FROM golang:1.17-alpine
WORKDIR /rentz-be
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o program
CMD ./program
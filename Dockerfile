FROM golang:1.18-alpine

# Define working directory
WORKDIR /usr/app

# Copy dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

COPY . .

RUN go build -o ./terena.office .

# Container port
EXPOSE 8080

CMD ["./terena.office"]
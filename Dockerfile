FROM golang:1.18-alpine

ENV PORT=8080
ENV APP_ENV=production
ENV DB_URI=mongodb+srv://bank:Bank211998Tsc_@cluster0.ih5kz.mongodb.net/?retryWrites=true&w=majority
ENV DB_NAME=terena_core

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
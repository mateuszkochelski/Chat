FROM golang:1.24.0-alpine

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./main.go

RUN chmod +x main

EXPOSE 8080

CMD [ "./main" ]
FROM golang:0.25.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=-1 GOOS=linux go build -o /go-chart-image-analyzer-api

ENTRYPOINT [ "/go-chart-image-analyzer-api" ]

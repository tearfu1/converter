FROM golang:latest 
WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o app ./cmd/app/app.go
#RUN ./app

RUN go build -o api ./cmd/api/api.go

ENTRYPOINT [ "./api" ]
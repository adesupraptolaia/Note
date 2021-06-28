# ambil images dari registry
# FROM [images_di_registry]
FROM golang:1.14

# COPY [file_local] [lokasi_di_docker]
COPY main.go /app/main.go
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

# tempat dimana RUN dan CMD bekerja
WORKDIR /app

RUN go mod tidy

# perintah cli
CMD ["go", "run", "/app/main.go"]
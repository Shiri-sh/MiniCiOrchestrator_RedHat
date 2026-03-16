FROM golang:1.23-alpine AS builder

# הגדרת תיקיית עבודה בתוך הקונטיינר
WORKDIR /app

# העתקת קבצי התלויות והורדתן
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

#final image
FROM alpine:latest

WORKDIR /root/

# copying the binary from the builder stage to the final image
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
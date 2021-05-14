FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/h24-assignment /app/
COPY frontend/ /app/frontend
WORKDIR /app
CMD ["./h24-assignment"]
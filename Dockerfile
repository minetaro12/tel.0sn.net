FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /work
COPY . /work/
RUN go build -ldflags "-s -w" -trimpath -o main

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /work/main /app/main

EXPOSE 8023

CMD ["/app/main"]
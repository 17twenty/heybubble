FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.12
WORKDIR /app/
COPY --from=builder /app/app /bin
COPY --from=builder /app/partials /app/partials
COPY --from=builder /app/static /app/static
COPY --from=builder /app/templates /app/templates
EXPOSE 8080
ENTRYPOINT [ "/bin/app" ]
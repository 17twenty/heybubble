FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.12
COPY --from=builder /app/app /bin
COPY --from=builder /app/partials /bin/partials
COPY --from=builder /app/templates /bin/templates
COPY --from=builder /app/static /bin/static
ENTRYPOINT [ "/bin/app" ]
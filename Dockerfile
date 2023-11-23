FROM alpine:latest

WORKDIR /app

COPY ports .
COPY ports.json .

# Use a non-root user for running the app
RUN adduser -D nonroot
RUN chown -R nonroot:nonroot /app
USER nonroot

EXPOSE 8080

CMD ["./ports"]

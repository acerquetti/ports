FROM alpine:latest

WORKDIR /app

COPY ports .
COPY ports.json .

# Create nonroot user w/o password
RUN adduser -D nonroot
# ... change files ownership
RUN chown -R nonroot:nonroot /app
# ... and use it
USER nonroot

EXPOSE 8080

CMD ["./ports"]

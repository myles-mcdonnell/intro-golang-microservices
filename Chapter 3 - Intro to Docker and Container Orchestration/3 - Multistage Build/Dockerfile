############################
# STEP 1 build executable binary
############################
FROM golang as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 go build -mod=vendor -a -o ./echoserver

############################
# STEP 2 build a small image
############################
FROM scratch
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/echoserver /echoserver

CMD ["./echoserver"]

FROM golang:1.18 as builder
WORKDIR /go/src/github.com/thalessathler/twitterlike
COPY . .
RUN make build

# FROM golang:1.18-alpine
# WORKDIR /home/twitterlike
# RUN apk add --no-cache \
#     tzdata \
#     ca-certificates
# RUN adduser olist \
#     --disabled-password
#     USER olist
# COPY --from=builder /go/src/github.com/thalessathler/twitterlike .
CMD ["./twitterlike"]

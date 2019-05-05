FROM golang:alpine as builder
RUN apk update && apk add git && mkdir --parents /go/src/github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs && cd /go/src/github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs
WORKDIR /go/src/github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs
ADD . ./
#RUN go get github.com/99designs/gqlgen && go install github.com/99designs/gqlgen
WORKDIR /go/src/github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=builder /go/src/github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server /go/src/svc
WORKDIR /go/src/svc
ENV PORT=3000
EXPOSE 3000 8085
CMD ["./main"]
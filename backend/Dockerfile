FROM golang:alpine as builder 
RUN mkdir /build
ADD . /build
WORKDIR /build/cmd
RUN go test && go build -o main .

FROM alpine
#RUN adduser -S -D -H -h /app appuser 
#-D Don't assign password
#-H Don't create home directory
#-S
#-h 
#USER appuser
COPY --from=builder /build/cmd/main /app/
WORKDIR /app
EXPOSE 8081
CMD ./main

FROM golang:1.12.0-alpine3.9
RUN apk --no-cache add gcc g++ make git
RUN mkdir /appservice
ADD . /appservice
WORKDIR /appservice
RUN go build -o ../main server/server.go server/resolver.go server/generated.go
CMD [ "/main" ]
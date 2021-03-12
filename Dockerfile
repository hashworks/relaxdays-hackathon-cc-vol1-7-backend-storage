FROM golang:1.16 AS build
WORKDIR /go/src
COPY ./ ./

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o server .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/server ./
EXPOSE 8080/tcp
ENTRYPOINT ["./server"]

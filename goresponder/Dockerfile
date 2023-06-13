FROM registry.access.redhat.com/ubi8/go-toolset:1.19.9-2 as builder
COPY . .
RUN go build -o server
RUN ls
RUN pwd

FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
COPY --from=builder /opt/app-root/src/server /app
EXPOSE 8080
CMD ["./server"]

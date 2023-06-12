FROM registry.access.redhat.com/ubi8/go-toolset:1.20.5
COPY . .
CMD ["go", "build", "-o", "server"]

FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
COPY server /app
EXPOSE 8080
#CMD ["dnf", "install", "mysql-community-server"]
#CMD ["dnf", "install", "libc.so.6"]
CMD ["./server"]
#CMD ["ls"]

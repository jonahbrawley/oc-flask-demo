FROM registry.access.redhat.com/ubi8/ubi
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
COPY test-sql /app
EXPOSE 8080
CMD ["dnf", "install", "mysql-community-server"]
CMD ["./test-sql"]
#CMD ["ls"]

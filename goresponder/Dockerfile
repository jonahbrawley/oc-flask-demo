FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
COPY test-sql /app
EXPOSE 8080
CMD ["./test-sql"]
#CMD ["ls"]

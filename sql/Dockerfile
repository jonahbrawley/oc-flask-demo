FROM registry.hub.docker.com/library/mysql
WORKDIR /

# copies setup script
COPY setup.sql /docker-entrypoint-initdb.d
EXPOSE 3306
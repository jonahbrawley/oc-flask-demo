FROM registry.hub.docker.com/library/mysql
WORKDIR /
COPY setup.sql /
# setup datamodel
CMD ["export", "MYSQL_ROOT_PASSWORD=\"password\""]
CMD ["mysql", "-uroot", "-ppassword", "<", "\"setup.sql\""]

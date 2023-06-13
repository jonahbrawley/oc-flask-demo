FROM registry.hub.docker.com/library/mysql
WORKDIR /
COPY setup.sql /
# setup datamodel
CMD ["mysql", "-uroot", "-ppassword", "<", "/source.sql"]

FROM golang:1.20.5
WORKDIR /app

# make needed dirs
#RUN mkdir /app/templates
#RUN mkdir /app/static

# copy required files
COPY server.go /app/
#COPY static/ /app/static
#COPY requirements.txt /app
#COPY app.py /app

# expose port
EXPOSE 8080

# install go mysql dependencies
CMD ["go", "mod", "init", "test-sql"]
CMD ["go", "mod", "tidy"]
CMD ["go", "get", "-u", "github.com/go-sql-driver/mysql"]

# start go server
CMD ["go", "run", "server.go"]

# start flask
#RUN pip3 install -r ./requirements.txt
#RUN export FLASK_APP=app
#RUN export FLASK_ENV=development
#CMD ["flask", "run"]
#ENTRYPOINT [ "python" ]
#CMD [ "app.py" ]

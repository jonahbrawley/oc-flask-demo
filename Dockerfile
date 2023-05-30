FROM python:3.11.2
WORKDIR /app

# copy required files
COPY templates/ /app
COPY static/ /app
COPY requirements.txt /app
COPY app.py /app

# start flask
RUN pip3 install -r ./requirements.txt
RUN export FLASK_APP=app
RUN export FLASK_ENV=development
CMD ["flask", "run"]

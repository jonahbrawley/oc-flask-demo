FROM python:3.11.2
RUN pip3 install -r requirements.txt
RUN export FLASK_APP=app
RUN export FLASK_ENV=development
CMD ["flask", "run"]

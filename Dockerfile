FROM python:3.11.2
WORKDIR /app

# make needed dirs
RUN mkdir /app/templates
RUN mkdir /app/static

# copy required files
COPY templates/ /app/templates
COPY static/ /app/static
COPY requirements.txt /app
COPY app.py /app

# expose port
EXPOSE 5000

# start gunicorn server (flask)
RUN pip3 install -r ./requirements.txt
CMD ["gunicorn", "app:app"]

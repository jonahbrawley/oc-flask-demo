FROM python:3.11.2
WORKDIR /app
COPY app.py /app
CMD ["python3", "app.py"]

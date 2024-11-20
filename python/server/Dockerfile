FROM python:3.9-alpine

RUN mkdir /app
WORKDIR /app
ADD server.py /app/
RUN apk add --update --no-cache wget curl iproute2

EXPOSE 5000
CMD ["python", "/app/server.py"]
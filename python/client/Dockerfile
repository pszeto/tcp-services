FROM python:3.9-alpine

RUN mkdir /app
WORKDIR /app
ADD client.py /app/
RUN apk add --update --no-cache wget curl iproute2

CMD ["/bin/sleep", "infinity"]
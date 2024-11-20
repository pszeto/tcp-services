#/bin/sh
tagNumber=13
cd server
docker build -t pszeto/python-tcp-server -f ./Dockerfile . --platform linux/amd64 --no-cache
cd ../client
docker build -t pszeto/python-tcp-client -f ./Dockerfile . --platform linux/amd64 --no-cache
cd ../auto-client
docker build -t pszeto/python-tcp-auto-client -f ./Dockerfile . --platform linux/amd64 --no-cache

docker tag pszeto/python-tcp-server pszeto/python-tcp-server:v${tagNumber}
docker tag pszeto/python-tcp-client pszeto/python-tcp-client:v${tagNumber}
docker tag pszeto/python-tcp-auto-client pszeto/python-tcp-auto-client:v${tagNumber}

docker push pszeto/python-tcp-server
docker push pszeto/python-tcp-server:v${tagNumber}
docker push pszeto/python-tcp-client
docker push pszeto/python-tcp-client:v${tagNumber}
docker push pszeto/python-tcp-auto-client
docker push pszeto/python-tcp-auto-client:v${tagNumber}
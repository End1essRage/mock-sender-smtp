#defaults
port ?= 8081
env = ENV_POD
network ?= my_network

SMTP_URL  ?= smtp_listener:25
HOST ?= :8080


#для докера
d_connect:
#как обычный флаг d_connect container=1234
	docker exec -it ${container} /bin/bash

d_run:
	docker run --rm -p ${port}:8080 -e ENV=$(env) \
	-e SMTP_URL=${SMTP_URL} \
	-e HOST=${HOST} \
	--network ${network} \
	end1essrage/mock-sender-smtp:latest

d_build: 
	docker build -t end1essrage/mock-sender-smtp .

d_push:
	docker push end1essrage/mock-sender-smtp:${tag}

#для подмена
p_connect:
#как обычный флаг p_connect container=1234
	podman exec -it ${container} /bin/bash
	
p_run:
	podman run -p ${port}:8080 -e ENV=$(env) \
	-e SMTP_URL=${SMTP_URL} \
	-e HOST=${HOST} \
	--network ${network} \
	end1essrage/mock-sender-smtp:latest

p_build: 
	podman build -t end1essrage/mock-sender-smtp .

p_push:
	podman push end1essrage/mock-sender-smtp:${tag}
all: docker

docker:
	docker build -t mheers/prometheus-exporter-sine .

push:
	docker push mheers/prometheus-exporter-sine

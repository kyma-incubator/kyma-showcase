.PHONY: build
build:
	docker build -t image .
	touch .env
	echo "REDIS_URL=\nREDIS_PASSWORD=\nPORT=" > .env
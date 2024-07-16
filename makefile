SHELL := /bin/bash

img:
	docker build \
		-f Dockerfile \
		-t aaa \
		.
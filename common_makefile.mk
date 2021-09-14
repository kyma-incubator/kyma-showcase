# Generic makefile - contains most used targets
# Required variables:
# IMG - name of the image, with docker push repository, full path and tag
# KMS_KEY_URL - required only when the target build is meant to be on post-submit or release job

# Build the docker image
.PHONY: docker-build
docker-build:
	docker build . -t ${IMG}

# Push the docker image
.PHONY: docker-push
docker-push:
	docker push ${IMG}

# CI specified targets
.PHONY: ci-pr
ci-pr: docker-build docker-push

.PHONY: ci-main
ci-main: docker-build docker-push cosign

.PHONY: ci-release
ci-release: docker-build docker-push cosign

# Cosign signing
.PHONY: cosign
cosign:
	cosign sign -key ${KMS_KEY_URL} ${IMG}

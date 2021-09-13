# Generic makefile - contains most used targets
# Required variables:
# IMG - name of the image, with docker push repository, full path and tag
# IMG_DOCKER_TAG - name of the image, with docker push repository, full path and DOCKER_TAG to satisfy Protescan requirements
# KMS_KEY_URL - required only when the target build is meant to be on post-submit or release job

# Build the docker image
.PHONY: docker-build
docker-build:
	docker build . -t ${IMG}

# Push the docker image
.PHONY: docker-push
docker-push:
	docker push ${IMG}

# Push additional tag to the existing image
.PHONY: docker-push-pr
docker-push-pr:
	docker push ${IMG_DOCKER_TAG}

# add second tag to the image
.PHONY: ci-tag
ci-tag:
	docker tag ${IMG} ${IMG_DOCKER_TAG}

# CI specified targets
.PHONY: ci-pr
ci-pr: docker-build ci-tag docker-push docker-push-pr

.PHONY: ci-main
ci-main: docker-build docker-push cosign

.PHONY: ci-release
ci-release: docker-build docker-push cosign

# Cosign signing
.PHONY: cosign
cosign:
	cosign sign -key ${KMS_KEY_URL} ${IMG}

# names for the images
IMG_FRONTEND="kyma-showcase-frontend"
IMG_BACKEND="kyma-showcase-backend"


# Build the frontend docker image
.PHONY: docker-build-frontend
docker-build-frontend:
	docker build ./frontend -t ${IMG_FRONTEND}

# Build the backend docker image
.PHONY: docker-build-backend
docker-build-backend:
	docker build ./backend -t ${IMG_BACKEND}

# CI specified targets
.PHONY: ci-frontend
ci-frontend: docker-build-frontend

.PHONY: ci-backend
ci-backend: docker-build-backend

.PHONY: all
all: docker-build-frontend docker-build-backend

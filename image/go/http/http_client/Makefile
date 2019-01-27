IMAGE := alextanhongpin/http
VERSION := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VCS_URL := $(shell git config --get remote.origin.url)
VCS_REF := $(shell git rev-parse HEAD)

NAME := $(shell basename `git rev-parse --show-toplevel`)
VENDOR := $(shell whoami)

print:
	@echo VERSION=${VERSION} 
	@echo BUILD_DATE=${BUILD_DATE}
	@echo VCS_URL=${VCS_URL}
	@echo VCS_REF=${VCS_REF}
	@echo NAME=${NAME}
	@echo VENDOR=${VENDOR}

build:
	docker build \
	--build-arg VERSION="${VERSION}" \
	--build-arg BUILD_DATE="${BUILD_DATE}" \
	--build-arg VCS_URL="${VCS_URL}" \
	--build-arg VCS_REF="${VCS_REF}" \
	--build-arg NAME="${NAME}" \
	--build-arg VENDOR="${VENDOR}" \
	-t ${IMAGE} .

run:
	@docker run -d -p 8080:8080 ${IMAGE} ./app -service_uri=http://localhost:8080/hello

label:
	@docker inspect --format='{{json .Config.Labels}}' ${IMAGE} 

call:
	# Example on calling the service deployed through Docker Desktop.
	curl http://localhost:$(shell kubectl get svc http-client -ojsonpath='{.spec.ports[0].nodePort}')

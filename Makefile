
# docker login --username=xxxxxx registry.cn-shanghai.aliyuncs.com
include .env
#export $(shell sed 's/=.*//' .env)


DockerHubUser = haierkeys
DockerHubName = obsidian-image-api-gateway


# DockerHubName		=	$(shell basename "$(PWD)")
projectRootDir	=	$(shell pwd)


ReleaseTagPre = release-v
DevelopTagPre = develop-v

P_NAME = api
P_BIN = image-api


platform = $(shell uname -m)

ifeq ($(platform),arm64)
	buildCmd = build
else
	buildCmd = build
endif

# These are the values we want to pass for Version and BuildTime
# GitTag	= $(shell git describe --tags)
GitTag	= $(shell git describe --tags --abbrev=0)
BuildTime=$(shell date +%FT%T%z)


# Go parameters
goCmd	=	go


# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X global.GitTag=$(GitTag) -X global.BuildTime=$(BuildTime)"

goBuild	=	$(goCmd) $(buildCmd) ${LDFLAGS}
goRun	=	$(goCmd) run ${LDFLAGS}

goClean	=	$(goCmd) clean
goTest	=	$(goCmd) test
goGet	=	$(goCmd) get -u



sourceDir	=	$(projectRootDir)
cfgDir		=	$(projectRootDir)/config
cfgFile		=	$(cfgDir)/config.yaml
buildDir	=	$(projectRootDir)/build


.PHONY: all build-all run test clean push-online push-dev build-macos-amd64 build-macos-arm64 build-linux-amd64 build-linux-arm64 build-winmdows-amd64
all: test build-all


build-all:
#	$(call checkStatic)
	$(MAKE) build-macos-amd64
	$(MAKE) build-macos-arm64
	$(MAKE) build-linux-amd64
	$(MAKE) build-linux-arm64
	$(MAKE) build-winmdows-amd64


run:
#	$(call checkStatic)
	$(call init)
	$(goRun)-v $(sourceDir)

# build2:
# 	$(call init)
# 	$(goBuild) -o $(binAdm) -v $(sourceAdmDir)
# 	$(goBuild) -o $(binNode) -v $(sourceNodeDir)
# 	mv $(binAdm) $(buildAdmDir)
# 	mv $(binNode) $(buildNodeDir)

test:
	@echo $(DockerHubName)
	@echo "Test Completed"

# $(goTest) -v -race -coverprofile=coverage.txt -covermode=atomic $(sourceAdmDir)
# $(goTest) -v -race -coverprofile=coverage.txt -covermode=atomic $(sourceNodeDir)
clean:
	rm -rf $(buildDir)

push-online:  build-linux
	$(call dockerImageClean)
	docker build --platform linux/amd64  -t  $(DockerHubUser)/$(DockerHubName):latest -f Dockerfile .
	docker tag  $(DockerHubUser)/$(DockerHubName):latest $(DockerHubUser)/$(DockerHubName):$(ReleaseTagPre)$(GitTag)

	docker push $(DockerHubUser)/$(DockerHubName):$(ReleaseTagPre)$(GitTag)
	docker push $(DockerHubUser)/$(DockerHubName):latest


push-dev:  build-linux
	$(call dockerImageClean)
	docker build --platform linux/amd64 -t $(DockerHubUser)/$(DockerHubName):dev-latest -f Dockerfile .
	docker tag $(DockerHubUser)/$(DockerHubName):dev-latest $(DockerHubUser)/$(DockerHubName):$(DevelopTagPre)$(GitTag)

	docker push $(DockerHubUser)/$(DockerHubName):$(DevelopTagPre)$(GitTag)
	docker push $(DockerHubUser)/$(DockerHubName):dev-latest



build-macos-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(goBuild) -o $(buildDir)/darwin_amd64/${P_BIN} $(bin) -v $(sourceDir)
build-macos-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(goBuild) -o $(buildDir)/darwin_arm64/${P_BIN} -v $(sourceDir)
build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(goBuild) -o $(buildDir)/linux_amd64/${P_BIN} -v $(sourceDir)
build-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(goBuild) -o $(buildDir)/linux_arm64/${P_BIN} -v $(sourceDir)
build-windows-amd64:
# CGO_ENABLED=0 CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp" $(goBuild) -o $(bin).exe -v $(sourceDir)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(goBuild) -o $(buildDir)/windows_amd64/${P_BIN}.exe -v $(sourceDir)
gox-linux:
	gox -osarch="linux/amd64 linux/arm64" -output="$(buildDir)/{{.OS}}_{{.Arch}}/${P_BIN}" ${LDFLAGS}
gox-all:
	gox -osarch="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64" -output="$(buildDir)/{{.OS}}_{{.Arch}}/${P_BIN}" ${LDFLAGS}

define dockerImageClean
	@echo "docker Image Clean"
	bash docker_image_clean.sh
endef

define init
	@echo "Build Init"
endef



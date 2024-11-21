
# docker login --username=xxxxxx registry.cn-shanghai.aliyuncs.com
include .env
#export $(shell sed 's/=.*//' .env)


RemoteDockerHub = haierkeys

ReleaseTagPre = release-v.
DevelopTagPre = develop-v.


platform = $(shell uname -m)

ifeq ($(platform),arm64)
	buildCmd = build
else
	buildCmd = build
endif

# These are the values we want to pass for Version and BuildTime
GitTag	= $(shell git describe --tags)
$(shell echo `git describe --abbrev=0`>VERSION)
BuildTime=$(shell date +%FT%T%z)


# Go parameters
goCmd	=	go
version	=	$(shell cat VERSION)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X global.GitTag=$(GitTag) -X global.BuildTime=$(BuildTime)"

goBuild	=	$(goCmd) $(buildCmd) ${LDFLAGS}
goRun	=	$(goCmd) run ${LDFLAGS}

goClean	=	$(goCmd) clean
goTest	=	$(goCmd) test
goGet	=	$(goCmd) get -u

projectName		=	$(shell basename "$(PWD)")
projectRootDir	=	$(shell pwd)


sourceDir	=	$(projectRootDir)
bin			=	$(projectRootDir)/image-api
cfgDir		=	$(projectRootDir)/config
cfgFile		=	$(cfgDir)/config.yaml
buildDir	=	$(projectRootDir)/build
buildCfgDir	=	$(buildDir)/config

.PHONY: all build build-all zip-all run test clean build-linux build-windows build-macos push-online push-dev build-t
all: test build
build:
#	$(call checkStatic)
	$(call init)
	$(goBuild) -o $(bin) -v $(sourceDir)
	@echo "Build OK"
	mv $(bin) $(buildDir)
	mkdir -p $(buildDir)/config
	cp $(cfgDir)/config.yaml $(buildDir)/config
	mkdir -p $(buildDir)/storage/logs
	mkdir -p $(buildDir)/storage/uploads
	mkdir -p $(buildDir)/storage/temp

build-all:
#	$(call checkStatic)
	$(MAKE) build-macos
	$(MAKE) build-linux
	$(MAKE) build-windows


publish-all: build-all
	tar -czvf $(buildDir)/image-api-windows-$(GitTag).tar.gz -C $(buildDir)/windows/ .
	tar -czvf $(buildDir)/image-api-macos-$(GitTag).tar.gz -C $(buildDir)/macos/ .
	tar -czvf $(buildDir)/image-api-linux-$(GitTag).tar.gz -C $(buildDir)/linux/ .

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
	@echo $(buildCmd)
	@echo "Test Completed"
# $(goTest) -v -race -coverprofile=coverage.txt -covermode=atomic $(sourceAdmDir)
# $(goTest) -v -race -coverprofile=coverage.txt -covermode=atomic $(sourceNodeDir)
clean:
	rm -rf $(buildDir)

push-online:  build-linux
	$(call dockerImageClean)
	docker build --platform linux/amd64  -t  $(RemoteDockerHub)/$(projectName):latest -f Dockerfile .
	docker tag  $(RemoteDockerHub)/$(projectName):latest $(RemoteDockerHub)/$(projectName):$(ReleaseTagPre)$(GitTag)

	docker push $(RemoteDockerHub)/$(projectName):$(ReleaseTagPre)$(GitTag)
	docker push $(RemoteDockerHub)/$(projectName):latest


push-dev:  build-linux
	$(call dockerImageClean)
	docker build --platform linux/amd64 -t $(RemoteDockerHub)/$(projectName):dev-latest -f Dockerfile .
	docker tag $(RemoteDockerHub)/$(projectName):dev-latest $(RemoteDockerHub)/$(projectName):$(DevelopTagPre)$(GitTag)

	docker push $(RemoteDockerHub)/$(projectName):$(DevelopTagPre)$(GitTag)
	docker push $(RemoteDockerHub)/$(projectName):dev-latest


build-macos:
	$(call init)
	rm -rf $(buildDir)/macos
	mkdir -p $(buildDir)/macos/
	GOOS=darwin GOARCH=amd64 $(goBuild) -o $(bin) -v $(sourceDir)
	mv $(bin) $(buildDir)/macos/
	mkdir -p $(buildDir)/macos/config
	cp $(cfgDir)/config.yaml $(buildDir)/macos/config
	mkdir -p $(buildDir)/macos/storage/logs
	mkdir -p $(buildDir)/macos/storage/uploads
	mkdir -p $(buildDir)/linux/storage/temp

build-linux:
	$(call init)
	rm -rf $(buildDir)/linux
	mkdir -p $(buildDir)/linux/
	GOOS=linux GOARCH=amd64 $(goBuild) -o $(bin) -v $(sourceDir)
	mv $(bin) $(buildDir)/linux/
	mkdir -p $(buildDir)/linux/config
	cp $(cfgDir)/config.yaml $(buildDir)/linux/config
	mkdir -p $(buildDir)/linux/storage/logs
	mkdir -p $(buildDir)/linux/storage/uploads
	mkdir -p $(buildDir)/linux/storage/temp


build-windows:
	$(call init)
	rm -rf $(buildDir)/windows
	mkdir -p $(buildDir)/windows/
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp" $(goBuild) -o $(bin).exe -v $(sourceDir)
	mv $(bin).exe $(buildDir)/windows/
	mkdir -p $(buildDir)/windows/config
	cp $(cfgDir)/config.yaml $(buildDir)/windows/config
	mkdir -p $(buildDir)/windows/storage/logs
	mkdir -p $(buildDir)/windows/storage/uploads
	mkdir -p $(buildDir)/linux/storage/temp



define dockerImageClean
	@echo "docker Image Clean"
	bash docker_image_clean.sh
endef

define init
	@echo "Build Init"
endef
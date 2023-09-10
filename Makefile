
# docker login --username=xxxxxx registry.cn-shanghai.aliyuncs.com
include .env
#export $(shell sed 's/=.*//' .env)


RemoteDockerHub = registry.cn-shanghai.aliyuncs.com/xxxxx/

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
BuildTime=$(shell date +%FT%T%z)


# Go parameters
goCmd	=	go
version	=	$(shell cat VERSION)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.GitTag=$(GitTag) -X main.BuildTime=$(BuildTime) -X main.Version=$(version)"

goBuild	=	$(goCmd) $(buildCmd) ${LDFLAGS}
goRun	=	$(goCmd) run ${LDFLAGS}

goClean	=	$(goCmd) clean
goTest	=	$(goCmd) test
goGet	=	$(goCmd) get -u

projectName		=	$(shell basename "$(PWD)")
projectRootDir	=	$(shell pwd)




sourceDir	=	$(projectRootDir)
bin			=	$(projectRootDir)/apiRun
cfgDir		=	$(projectRootDir)/configs
cfgFile		=	$(cfgDir)/config.yaml
keystoreDir	=	$(projectRootDir)/keystore
buildDir	=	$(projectRootDir)/build
buildCfgDir	=	$(buildDir)/configs

.PHONY: all build run test clean build-linux build-windows build-macos push-online push-dev build-t
all: test build
build:
#	$(call checkStatic)
	$(call init)
	$(goBuild) -o $(bin) -v $(sourceDir)
	@echo "Build OK"
	mv $(bin) $(buildDir)
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
	docker build -t  $(RemoteDockerHub)$(projectName):online-latest -f Dockerfile .
#	docker tag $(projectName):$(ReleaseTagPre)$(GitTag) $(projectName):latest
	docker tag  $(RemoteDockerHub)$(projectName):online-latest $(RemoteDockerHub)$(projectName):$(ReleaseTagPre)$(GitTag)

	docker push $(RemoteDockerHub)$(projectName):$(ReleaseTagPre)$(GitTag)
	docker push $(RemoteDockerHub)$(projectName):online-latest

push-dev:  build-linux
	$(call dockerImageClean)
	docker build -t $(RemoteDockerHub)$(projectName):dev-latest -f Dockerfile .
	docker tag $(RemoteDockerHub)$(projectName):dev-latest $(RemoteDockerHub)$(projectName):$(DevelopTagPre)$(GitTag)

	docker push $(RemoteDockerHub)$(projectName):$(DevelopTagPre)$(GitTag)
	docker push $(RemoteDockerHub)$(projectName):dev-latest


build-macos:
	$(call init)
	GOOS=darwin GOARCH=amd64 $(goBuild) -o $(bin) -v $(sourceDir)
	mv $(bin) $(buildDir)

build-linux:
	$(call init)
	GOOS=linux GOARCH=amd64 $(goBuild) -o $(bin) -v $(sourceDir)
	mv $(bin) $(buildDir)

build-windows:
	$(call init)
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp" $(goBuild) -o $(bin).exe -v $(sourceDir)
	mv $(bin).exe $(buildDir)

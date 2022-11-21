OUTPUT_PATH := ./build
CMD_NAME := myprobe
DATE_NOW := $(shell date "+%Y%m%d_%H%M%S")

.PHONY: clean start build_linux build_mac build_windows end docker
all: clean start build_linux build_mac build_windows end


#清理开发目录
clean:
	rm -rf ./build/*
	rm -rf ./*.log
	rm -rf ./*.tar.gz


#整理开发依赖
start:
	go mod tidy -compat=1.17


#编译为Linux发行版
build_linux:
	mkdir -p "${OUTPUT_PATH}/linux"
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=0 \
	go build -ldflags '-extldflags "-static"' -o "${OUTPUT_PATH}/linux/${CMD_NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/linux/"
	tar -zcvf ${OUTPUT_PATH}/myprobe-linux-amd64.tar.gz -C ${OUTPUT_PATH} linux


#编译为Mac发行版
build_mac:
	mkdir -p "${OUTPUT_PATH}/mac"
	GOARCH=amd64 \
	GOOS=darwin \
	CGO_ENABLED=0 \
	go build -ldflags '-extldflags "-static"' -o "${OUTPUT_PATH}/mac/${CMD_NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/mac/"
	tar -zcvf ${OUTPUT_PATH}/myprobe-mac-amd64.tar.gz -C ${OUTPUT_PATH} mac

#编译为windows发行版
build_windows:
	mkdir -p "${OUTPUT_PATH}/windows"
	GOARCH=amd64 \
	GOOS=darwin \
	CGO_ENABLED=0 \
	go build -ldflags '-extldflags "-static"' -o "${OUTPUT_PATH}/windows/${CMD_NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/windows/"
	tar -zcvf ${OUTPUT_PATH}/myprobe-windows-amd64.tar.gz -C ${OUTPUT_PATH} windows


#结束之前做些什么
end:
	ls -alh ${OUTPUT_PATH}

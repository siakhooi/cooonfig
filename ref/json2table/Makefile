clean:
	rm -rf bin
go-mod-tidy:
	go mod tidy
build:
	scripts/build.sh
	goreleaser build --snapshot --clean
go-release:
	goreleaser release --snapshot --clean
build-linux:
	scripts/build.sh -l 2>&1 | tee build-linux.log
test:
	scripts/test.sh
golangci-lint:
	golangci-lint run 2>&1 | tee golangci-lint.log

deb-list:
	 dpkg-deb --contents dist/siakhooi-json2table*_amd64.deb
all: clean go-mod-tidy test golangci-lint go-release
quick: clean go-mod-tidy test golangci-lint build-linux

release:
	scripts/create-release.sh

commit-watch: commit
	gh run watch

release-watch: release
	gh run watch

run:
	bin/json2table-linux-amd64
run-help:
	bin/json2table-linux-amd64 -h
run-version:
	bin/json2table-linux-amd64 -v
run-build:
	bin/json2table-linux-amd64 --build
run-no-arguments-1:
	bin/json2table-linux-amd64
run-no-arguments-2:
	bin/json2table-linux-amd64 -s ./samples/spec1.json
run-too-many-arguments-1:
	bin/json2table-linux-amd64 ./samples/data1.json ./samples/data1.json
run-too-many-arguments-2:
	bin/json2table-linux-amd64 -s ./samples/spec.json ./samples/data1.json ./samples/data1.json
run-1:
	bin/json2table-linux-amd64 -s ./samples/spec1.json ./samples/data1.json
run-1a:
	bin/json2table-linux-amd64 -s ./samples/spec2.json ./samples/data1.json
run-1c:
	bin/json2table-linux-amd64 -c 'id,desc,url,display.name' ./samples/data2.json
run-2:
	bin/json2table-linux-amd64 --spec ./samples/spec1.json ./samples/data1.json
run-3:
	cat ./samples/data1.json |bin/json2table-linux-amd64 -s ./samples/spec1.json
run-4:
	JSON2TABLE_SPEC_FILE=./samples/spec1.json bin/json2table-linux-amd64 ./samples/data1.json
run-5:
	cat  ./samples/data1.json |JSON2TABLE_SPEC_FILE=./samples/spec1.json bin/json2table-linux-amd64
run-6:
	JSON2TABLE_SPEC='{"dataPath":"$.data2","columns":[{"path":"id","title":"ID"},{"path":"display.name"}]}' bin/json2table-linux-amd64 ./samples/data1.json

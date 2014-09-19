SHELL := /bin/bash
TARGETS = wikicats wikinorm wikidatatojson wikitojson

# http://docs.travis-ci.com/user/languages/go/#Default-Test-Script
test:
	go test -v

bench:
	go test -bench=.

imports:
	goimports -w .

fmt:
	go fmt ./...

vet:
	go vet ./...

all: fmt test
	go build

install:
	go install

clean:
	go clean
	rm -f coverage.out
	rm -f $(TARGETS)
	rm -f wikitools-*.x86_64.rpm
	rm -f debian/wikitools*.deb
	rm -rf debian/wikitools/usr

cover:
	go get -d && go test -v	-coverprofile=coverage.out
	go tool cover -html=coverage.out

wikicats:
	go build cmd/wikicats/wikicats.go

wikinorm:
	go build cmd/wikinorm/wikinorm.go

wikidatatojson:
	go build cmd/wikidatatojson/wikidatatojson.go

wikitojson:
	go build cmd/wikitojson/wikitojson.go

# ==== packaging

deb: $(TARGETS)
	mkdir -p debian/wikitools/usr/sbin
	cp $(TARGETS) debian/wikitools/usr/sbin
	cd debian && fakeroot dpkg-deb --build wikitools .

REPOPATH = /usr/share/nginx/html/repo/CentOS/6/x86_64

publish: rpm
	cp wikitools-*.rpm $(REPOPATH)
	createrepo $(REPOPATH)

rpm: $(TARGETS)
	mkdir -p $(HOME)/rpmbuild/{BUILD,SOURCES,SPECS,RPMS}
	cp ./packaging/wikitools.spec $(HOME)/rpmbuild/SPECS
	cp $(TARGETS) $(HOME)/rpmbuild/BUILD
	./packaging/buildrpm.sh wikitools
	cp $(HOME)/rpmbuild/RPMS/x86_64/wikitools*.rpm .

.PHONY: all clean

all: dist.zip

dist.zip: 10chapters public/* tmpl/*
	zip $@ $^

10chapters: *.go
	go build

clean:
	-rm -rf dist.zip
	-rm -rf 10chapters

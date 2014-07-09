all:
	export GOPATH=`pwd` && go install all
init:
	git config --global user.email "go@hetao.name"
	git config --global user.name "user.name=go-study"
	cd src ; go get -d github.com/garyburd/redigo && go get -d github.com/gosexy/redis && go get -d github.com/qpliu/qrencode-go && go get -d menteslibres.net/gosexy/to

APPNAME=model_auto

export GOPATH=/Users/kevinchen/Documents/Golang/koalaone/


.PHONY : vgo
vgo:
	@echo "GOPATH:"${GOPATH}
	vgo build

.PHONY : fix
fix:
	@echo "GOPATH:"${GOPATH}
	vgo fix

.PHONY : vendor
vendor:
	@echo "GOPATH:"${GOPATH}
	vgo mod vendor


.PHONY : build
build:
	@echo "GOPATH:"${GOPATH}
	go build -o ${APPNAME}


.PHONY : install
install:
	@echo "GOPATH:"${GOPATH}
	vgo install
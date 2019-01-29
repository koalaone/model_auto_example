APPNAME=model_auto_example

export GOPATH=/Users/kevinchen/Documents/Golang/koalaone/


.PHONY : vgo
vgo:
	@echo "GOPATH:"${GOPATH}
	vgo get -u
	vgo clean
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


.PHONY : run
run:
	@echo "GOPATH:"${GOPATH}
	vgo run main.go
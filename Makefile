.PHONY: build
build:

.PHONY: install
install:
	go build -o  ./build/test-cni
	sudo cp build/test-cni /opt/cni/bin/
	sudo chmod +x /opt/cni/bin/test-cni

.PHONY: tail.log
tail.log:
	tail -f /var/log/test-cni.log
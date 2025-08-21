exec=health.exe
syso=health.syso
build=./build.sh

build: 
	$(build)

delete:
	rm -fr $(exec)
	rm -fr $(syso)
.PHONY: clean profile

all: gobasic gothreaded cbasic cmmap

profile: gen all
	time ./go-basic/main.out ./gen/behemoth_00p_0_bw.raw
	time ./go-threaded/main.out ./gen/behemoth_00p_0_bw.raw
	time ./cpp-basic/main.out ./gen/behemoth_00p_0_bw.raw
	time ./cpp-mmap/main.out ./gen/behemoth_00p_0_bw.raw

gen:
	go generate ./...

gobasic:
	(cd go-basic && make)

gothreaded:
	(cd go-threaded && make)

cbasic:
	(cd cpp-basic && make)

cmmap:
	(cd cpp-mmap && make)

clean:
	(cd go-basic && make clean)
	(cd go-threaded && make clean)
	(cd cpp-basic && make clean)
	(cd cpp-mmap && make clean)

#!/bin/sh
#clang -MD -MF xconnect.d -target bpf -I ~/linux/tools/lib/bpf -c xconnect.c

#clang -target bpf -Wall -O2 -emit-llvm -g -Iinclude -c xconnect.c -o - | \
#llc -march=bpf -mcpu=probe -filetype=obj -o xconnect.o

clang -O2 -emit-llvm -c xconnect.c -o - | \
	llc -march=bpf -mcpu=probe -filetype=obj -o xconnect.o

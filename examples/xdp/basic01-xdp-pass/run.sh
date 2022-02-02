#!/bin/sh

bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
# clang -g -O2 -target bpf -D__TARGET_ARCH_x86_64 -I . -c xdp_pass_kern.c -o xdp_pass_kern.o

clang -O2 -emit-llvm -c xdp_pass_kern.c -o - | \
	llc -march=bpf -mcpu=probe -filetype=obj -o xdp_pass_kern.o

# bpftool gen skeleton hello.bpf.o > hello.skel.h
# clang -g -O2 -Wall -I . -c hello.c -o hello.o
# clang -Wall -O2 -g hello.o ../../libbpf/build/libbpf.a -lelf -lz -o hello
# sudo ./hello
#!/bin/sh

bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
clang -g -O2 -target bpf -D__TARGET_ARCH_x86_64 -I . -c hello.bpf.c -o hello.bpf.o
bpftool gen skeleton hello.bpf.o > hello.skel.h
clang -g -O2 -Wall -I . -c hello.c -o hello.o
clang -Wall -O2 -g hello.o ../../libbpf/build/libbpf.a -lelf -lz -o hello
sudo ./hello
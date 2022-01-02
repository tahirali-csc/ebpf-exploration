#!/bin/sh

bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
clang -g -O2 -target bpf -D__TARGET_ARCH_x86_64 -I . -c maps.bpf.c -o maps.bpf.o
bpftool gen skeleton maps.bpf.o > maps.skel.h
clang -g -O2 -Wall -I . -c maps.c -o maps.o
clang -Wall -O2 -g maps.o ../../libbpf/build/libbpf.a -lelf -lz -o maps
sudo ./maps
Compiling BPF C
--------
We rely BPF C library to build the loader in C. https://github.com/libbpf/libbpf

```bash
cd src
make BUILD_STATIC_ONLY=1 OBJDIR=../build/libbpf DESTDIR=../build INCLUDEDIR= LIBDIR= UAPIDIR= install
```
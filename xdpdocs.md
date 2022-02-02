**XDP Tutorial**
------
1. https://github.com/xdp-project/xdp-tutorial. Quirks in compiling example.
```
https://github.com/libbpf/libbpf/issues/128
```

2. https://duo.com/labs/tech-notes/writing-an-xdp-network-filter-with-ebpf
3. https://developers.redhat.com/blog/2021/04/01/get-started-with-xdp#task_3__map_and_count_the_processed_packets
4. https://developers.redhat.com/blog/2021/04/01/get-started-with-xdp

*linux ip command cheat sheet*
https://access.redhat.com/sites/default/files/attachments/rh_ip_command_cheatsheet_1214_jcs_print.pdf

Introduction to Linux interfaces for virtual networking
https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking#

*Some issue in building old xpd-tutorial*
Add `-lz` in linking

+LIBS = -l:libbpf.a -lelf $(USER_LIBS) -lz
# ebpf-exploration

**Hello World example**
--------
Good start from here:
https://www.sartura.hr/blog/simple-ebpf-core-application/

The code is here
https://github.com/sartura/ebpf-hello-world

**BPF and XDP Reference Guide**
--------
1. https://docs.cilium.io/en/stable/bpf/


**Linux events by category:**
--------
```bash
/sys/kernel/debug/tracing/events
```
for example, to list all sys events
```bash
sudo ls -la /sys/kernel/debug/tracing/events/syscalls
```

**bcc Reference Guide**
--------
https://github.com/iovisor/bcc/blob/master/docs/reference_guide.md#6-bpf_get_current_comm

**XDP Tutorial**
------
1. https://github.com/xdp-project/xdp-tutorial. Quirks in compiling example.
```
https://github.com/libbpf/libbpf/issues/128
```

2. https://duo.com/labs/tech-notes/writing-an-xdp-network-filter-with-ebpf
3. https://developers.redhat.com/blog/2021/04/01/get-started-with-xdp#task_3__map_and_count_the_processed_packets

*linux ip command cheat sheet*
https://access.redhat.com/sites/default/files/attachments/rh_ip_command_cheatsheet_1214_jcs_print.pdf
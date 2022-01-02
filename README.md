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
2. https://duo.com/labs/tech-notes/writing-an-xdp-network-filter-with-ebpf
3. https://developers.redhat.com/blog/2021/04/01/get-started-with-xdp#task_3__map_and_count_the_processed_packets

**Linux events by category:**
--------
```bash
/sys/kernel/debug/tracing/events
```
for example, to list all sys events
```bash
sudo ls -la /sys/kernel/debug/tracing/events/syscalls
```



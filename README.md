# ebpf-exploration

**Hello World example**
--------
Good start from here:
https://www.sartura.hr/blog/simple-ebpf-core-application/

The code is here
https://github.com/sartura/ebpf-hello-world

**BPF and XDP Reference Guide**
https://docs.cilium.io/en/stable/bpf/

**Linux events by category:**
--------
```bash
/sys/kernel/debug/tracing/events
```
for example, to list all sys events
```bash
sudo ls -la /sys/kernel/debug/tracing/events/syscalls
```

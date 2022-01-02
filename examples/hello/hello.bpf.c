#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

SEC("tracepoint/syscalls/sys_enter_execve")
int tracepoint__syscalls__sys_enter_execve(struct trace_event_raw_sys_enter *args)
{
	//the program first attaches to the execve system call
	//prints formatted output to the system's trace pipe in /sys/kernel/debug/tracing/trace_pipe
	bpf_printk("Hello world!\n");
	return 0;
}

char LICENSE[] SEC("license") = "GPL";

#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include "maps.h"

/*
Defines a hash map that works like a hash table. It has a maximum of 128 entries, 
where the key is the PID, and the value is a struct that will track execve arguments. 
When defining maps, it is important to place them in the ".maps" section so that 
the userspace loader can load and set up the maps properly. 

https://man7.org/linux/man-pages/man2/bpf.2.html
*/
struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, 128);
	//pid_t is defined in vmlinux.h as "typedef __kernel_pid_t pid_t";
	__type(key, pid_t);
	__type(value, struct event);
} execs SEC(".maps");

SEC("tracepoint/syscalls/sys_enter_execve")
int tracepoint__syscalls__sys_enter_execve(struct trace_event_raw_sys_enter *ctx)
{
	struct event *event;
	pid_t pid;
	u64 id;
	uid_t uid = (u32) bpf_get_current_uid_gid();

	id = bpf_get_current_pid_tgid();
	pid = (pid_t)id;

	//Create or update an element (key/value pair) in a specified map.
	if (bpf_map_update_elem(&execs, &pid, &((struct event){}), 1)) {
		return 0;
	}

	//Look up an element by key in a specified map and return its value
	event = bpf_map_lookup_elem(&execs, &pid);
	if (!event) {
		return 0;
	}

	//Update event object
	event->pid = pid;
	event->uid = uid;
	//Populates the first argument address with the current process name.
	bpf_get_current_comm(&event->comm, sizeof(event->comm));

	/*
	if (bpf_map_update_elem(&execs, &pid, event, 2)) {
		return 0;
	}
	*/

	return 0;
}

char LICENSE[] SEC("license") = "GPL";
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

SEC("xdp")
int  xdp_xconnect(struct xdp_md *ctx)
{
  return 1;
}

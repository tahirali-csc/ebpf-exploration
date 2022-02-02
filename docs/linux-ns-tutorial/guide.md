Install vSwitch on Centos8
---------------------
https://www.osradar.com/how-to-install-open-vswitch-on-centos-8-rhel-8/

*Linux net link handy commands*

Turn off program(xdp) on interface
---------------------
```bash
ip link set veth-basic02 xdp off
```

Delete network interface
---------------------
```bash
ip link delete dev veth-basic03
```
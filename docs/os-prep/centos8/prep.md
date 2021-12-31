Centos 8 -- Dev Environment
--------------

**Verify kernel options**
--------------------
Make sure ```CONFIG_DEBUG_INFO_BTF``` and ```CONFIG_DEBUG_INFO``` are set.

```bash
cat /boot/config-`uname -r` | grep 'CONFIG_DEBUG_INFO_BTF'
```

In case if you are not updated, follow this:
https://linuxconfig.org/how-to-update-centos

**libbpf installation**
--------------------
```sh
sudo yum install libbpf
```

**Download libbpf header files**
--------------------
https://centos.pkgs.org/8-stream/centos-powertools-x86_64/libbpf-devel-0.2.0-1.el8.x86_64.rpm.html

```bash
sudo dnf --enablerepo=powertools install libbpf-devel
```

You may runt into ```Error: Unknown repo: 'powertools'```

Got some ideas from here: https://serverfault.com/questions/997896/how-to-enable-powertools-repository-in-centos-8

**ssh configuration**(optional for dev testing)
--------------------
https://linuxconfig.org/redhat-8-enable-ssh-service

allow ssh
```sh
firewall-cmd --zone=public --permanent --add-service=ssh
```

*Network quirks on Virtual Box*
--------------------
Make sure to connect to right network when sshing from the host machine. Choose the matching network from top list(hierarchy icon).
### Prepare rootfs

```
https://wiki.alpinelinux.org/wiki/Installing_Alpine_Linux_in_a_chroot

wget http://mirrors.sjtug.sjtu.edu.cn/alpine/latest-stable/main/x86_64/apk-tools-static-2.10.1-r0.apk

./sbin/apk.static -X http://mirrors.sjtug.sjtu.edu.cn/alpine/latest-stable/main -U --allow-untrusted --root /rootfs --initdb add alpine-base

```

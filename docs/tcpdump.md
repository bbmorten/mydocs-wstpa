#  tcpdump samples

## Example for rotating files

a tcpdump command that listens on a specific interface (eth0) for 5 minutes, creates a new file every 30 seconds with a timestamp in the filename, and only keeps the 5 most recent files:

```shell
tcpdump -i eth0 -G 30 -W 5 -w tcpdump-%Y%m%d%H%M%S.pcap -Z root -G 300
```

MACOSx version needs editing

```shell
while sleep 30; do tcpdump -i eth0 -w tcpdump-$(date +%Y%m%d%H%M%S).pcap -Z root -W 5; done
```

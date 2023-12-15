
# same simple trafic analizer
```
go run main.go -pcap date.pcapng -filter "tcp and port 443 and host 10.10.0.136"
```

## работает на основе BPF filter
```
https://www.ibm.com/docs/en/qsip/7.4?topic=queries-berkeley-packet-filters
```

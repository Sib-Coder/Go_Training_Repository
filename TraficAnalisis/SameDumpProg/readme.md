# TCPDump для тех кто не умеет его юзать
# start
```
go run main.go --fname "dump.pcap" --iface "wlp3s0"
```
# пример использования
```
[root@fedora awesomeProject1]# go run main.go --fname "dump.pcap" --iface "wlp3s0"
2023/12/16 12:17:05 INFO Start Program Trafik Analisis
Listening for traffic...
Packet captured and written to pcap file
Packet captured and written to pcap file
Packet captured and written to pcap file
Packet captured and written to pcap file
Packet captured and written to pcap file
Packet captured and written to pcap file
...
```
package main

import (
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log/slog"
)

var filename = flag.String("pcap", "", "Pcap File to Load and Parse")
var filter = flag.String("filter", "", "BPF Filter to apply")

func main() {
	slog.Info("Start Program Trafik Analisis")
	defer slog.Info("End Program")

	flag.Parse()

	if *filename == "" {
		flag.PrintDefaults()
		slog.Error("no pcap file parameter was passed")
	}

	pcapHandle, err := pcap.OpenOffline(*filename)
	if err != nil {
		slog.Error("error opening file - %v", err)
	}

	defer pcapHandle.Close()

	if *filter != "" {
		err = pcapHandle.SetBPFFilter(*filter)
		if err != nil {
			slog.Error("error appling filter %s -%v", *filter, err)
		}
	}

	packetsFiltered := gopacket.NewPacketSource(pcapHandle, pcapHandle.LinkType())
	for packet := range packetsFiltered.Packets() {
		fmt.Println(packet)
	}
}


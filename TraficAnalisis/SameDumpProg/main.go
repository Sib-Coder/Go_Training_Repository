package main

import (
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"log"
	"log/slog"
	"os"
)

var fname = flag.String("fname", "", "Give name file to damp")
var iface = flag.String("iface", "", "Give interface name to damp") //wlp3s0

func main() {
	slog.Info("Start Program Trafik Analisis")
	defer slog.Info("End Program")

	flag.Parse()

	if *fname == "" {
		flag.PrintDefaults()
		slog.Error("no pcap file parameter was passed")
	}
	if *iface == "" {
		flag.PrintDefaults()
		slog.Error("no interface name parameter was passed")
	}

	ProgrammStart(*iface, *fname)
}

func ProgrammStart(iface string, fname string) {
	//открываем интерфасе на запись
	handle, err := pcap.OpenLive(iface, 6553, true, pcap.BlockForever)
	if err != nil {
		slog.Error("error open iface", err)
	}
	defer handle.Close()

	//открытие файла
	file, err := os.Create(fname)
	if err != nil {
		slog.Error("errror open file", err)
	}
	defer file.Close()

	// Создаем писателя для pcap файла
	writer := pcapgo.NewWriter(file)
	
	// Записываем заголовок pcap файла
	err = writer.WriteFileHeader(65536, layers.LinkTypeEthernet)
	if err != nil {
		slog.Error("writer error", err)
	}
	fmt.Println("Listening for traffic...")

	// Бесконечный цикл для прослушивания пакетов сетевого трафика
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Записываем пакет в файл pcap
		err := writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Packet captured and written to pcap file")
	}
}


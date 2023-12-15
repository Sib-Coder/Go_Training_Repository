package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"math/rand"
	"strconv"
)

func main() {
	// Open the pcap file
	handle, err := pcap.OpenOffline("date.pcapng")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Create a packet source to decode packets from the pcap file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Iterate through each packet in the pcap file
	for packet := range packetSource.Packets() {
		fmt.Println("new packet")
		//Add Packet in DB
		IdPacket, err := AddPacketDataInDB(packet)
		if err != nil {
			fmt.Println("Problem add to database")
		}
		fmt.Println("New ID is:" + strconv.Itoa(IdPacket))

		//Add id in Kafka
		err = NotificationKafka(IdPacket)
		if err != nil {
			fmt.Println("Problem add to Kafka")
		}

	}
}

// /Затычки будующих интерфейсов
func AddPacketDataInDB(packet gopacket.Packet) (Id int, err error) {
	fmt.Println(packet.Data())
	return rand.Int(), nil
}

func NotificationKafka(id int) error {
	fmt.Println("id :" + strconv.Itoa(id) + " add to Kafka")
	return nil
}


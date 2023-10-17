package main

import (
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func getBandwidth(packet gopacket.Packet, startTime time.Time, totalBytes *int64) []interface{} {
	length := int64(len(packet.Data()))
	timestamp := packet.Metadata().Timestamp

	elapsedTime := timestamp.Sub(startTime).Seconds()
	bandwidthBps := float64(length*8) / elapsedTime
	*totalBytes += length
	log.Printf("TimeStamp : %v , BandWidth : %.2f bps", timestamp, bandwidthBps)
	var res []interface{}
	res = append(res, timestamp)
	res = append(res, bandwidthBps)
	return res
}
func main() {
	device := "en0"
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	startTime := time.Now()
	totalBytes := int64(0)
	for packet := range source.Packets() {
		getBandwidth(packet, startTime, &totalBytes)
	}

}

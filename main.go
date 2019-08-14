package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
	"log"
	"net"
	"os"
	"packetspy-server/utils"
	"strconv"
	"strings"
	"time"
)






func HandleClient(conn net.Conn) error {
	a := conn.RemoteAddr().String()
	addr := strings.Split(a, ":")[0]
	fname := addr + "__" + utils.PrintableDate(time.Now()) + ".pcap"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}

	bufrdr := bufio.NewReader(conn)
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(65535), layers.LinkTypeEthernet)
	defer f.Close()

	for {
		p, _ := bufrdr.ReadString('\n')
		pb, err := hex.DecodeString(strings.TrimSuffix(p, "\n"))
		if err != nil {
			log.Fatal(err)
		}
		if len(pb) > 0 {
			fields := strings.Split(string(pb), ",")
			if len(pb) < 4 {
				log.Fatal(errors.New("Too few fields received"))
			}
			pkt := gopacket.NewPacket([]byte(fields[4]), layers.LayerTypeEthernet, gopacket.Default)
			info := utils.MakeCaptureInfo(fields[:4])
			err := w.WritePacket(info, pkt.Data())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}



func main() {
	fmt.Println("starting...")
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Receiving...")
		go HandleClient(conn)
	}
}

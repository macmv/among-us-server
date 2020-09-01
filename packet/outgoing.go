package packet

import (
  "net"
  "fmt"
  "strings"
  "strconv"
)

type OutgoingPacket struct {
  data []byte
}

func NewOutgoingPacket() *OutgoingPacket {
  p := OutgoingPacket{}
  p.data = []byte{}
  return &p
}

func (o *OutgoingPacket) Send(conn *net.UDPConn, addr *net.UDPAddr) {
  fmt.Println("Sending data", o.data)
  conn.WriteToUDP(o.data, addr)
}

func (o *OutgoingPacket) WriteByte(val byte) {
  o.data = append(o.data, val)
}

func (o *OutgoingPacket) WriteString(val string) {
  o.data = append(o.data, byte(len(val)))
  o.data = append(o.data, []byte(val)...)
}

func (o *OutgoingPacket) WriteIP(val string) {
  sections := strings.Split(val, ".")
  if len(sections) != 4 {
    panic("Invalid ip!")
  }
  for _, s := range sections {
    val, err := strconv.ParseUint(s, 10, 8)
    if err != nil || val < 0 || val > 255 {
      panic(err)
      panic("Invalid number in the ip!")
    }
    o.WriteByte(byte(val))
  }
}

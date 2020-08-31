package packet

import (
  "net"
)

type OugoingPacket struct {
  data []byte
}

func NewOutgoingPacket() *OugoingPacket {
  p := OugoingPacket{}
  p.data = []byte{}
  return &p
}

func (o *OugoingPacket) Send(conn *net.UDPConn, addr *net.UDPAddr) {
  conn.WriteToUDP(o.data, addr)
}

func (o *OugoingPacket) WriteByte(val byte) {
  o.data = append(o.data, val)
}

func (o *OugoingPacket) WriteString(val string) {
  o.data = append(o.data, byte(len(val)))
  o.data = append(o.data, []byte(val)...)
}

package packet

import (
  "net"
  "math"
  "strings"
  "strconv"
  "encoding/binary"
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
  conn.WriteToUDP(o.data, addr)
}

func (o *OutgoingPacket) Data() []byte {
  return o.data
}

func (o *OutgoingPacket) WriteByte(val byte) {
  o.data = append(o.data, val)
}

func (o *OutgoingPacket) WriteBytes(val []byte) {
  o.data = append(o.data, val...)
}

func (o *OutgoingPacket) WriteInt(val int32) {
  bytes := make([]byte, 4)
  binary.LittleEndian.PutUint32(bytes, uint32(val))
  o.data = append(o.data, bytes...)
}

func (o *OutgoingPacket) WriteFloat(val float32) {
  bits := math.Float32bits(val)
  bytes := make([]byte, 4)
  binary.LittleEndian.PutUint32(bytes, bits)
  o.data = append(o.data, bytes...)
}

func (o *OutgoingPacket) WriteDouble(val float64) {
  bits := math.Float64bits(val)
  bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(bytes, bits)
  o.data = append(o.data, bytes...)
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

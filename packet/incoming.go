package packet

import (
  "math"
  "encoding/binary"
)

type IncomingPacket struct {
  data []byte
  index int
}

func NewIncomingPacketFromBytes(bytes []byte) *IncomingPacket {
  p := IncomingPacket{}
  p.data = bytes
  p.index = 1
  return &p
}

func (p *IncomingPacket) Id() byte {
  return p.data[0]
}

func (p *IncomingPacket) Remaining() []byte {
  return p.data[p.index:]
}

func (p *IncomingPacket) ReadByte() byte {
  val := p.data[p.index]
  p.index += 1
  return val
}

func (p *IncomingPacket) ReadShort() uint16 {
  val := binary.LittleEndian.Uint16(p.data[p.index:p.index+2])
  p.index += 2
  return val
}

func (p *IncomingPacket) ReadInt() uint32 {
  val := binary.LittleEndian.Uint32(p.data[p.index:p.index+4])
  p.index += 4
  return val
}

func (p *IncomingPacket) ReadLong() uint64 {
  val := binary.LittleEndian.Uint64(p.data[p.index:p.index+8])
  p.index += 8
  return val
}

func (p *IncomingPacket) ReadFloat() float32 {
  val := math.Float32frombits(binary.LittleEndian.Uint32(p.data[p.index:p.index+4]))
  p.index += 4
  return val
}

func (p *IncomingPacket) ReadDouble() float64 {
  val := math.Float64frombits(binary.LittleEndian.Uint64(p.data[p.index:p.index+8]))
  p.index += 8
  return val
}

func (p *IncomingPacket) ReadString() string {
  length := int(p.data[p.index] + 1)
  val := string(p.data[p.index+1:p.index+length])
  p.index += length
  return val
}

func (p *IncomingPacket) ReadBytes(num int) []byte {
  val := p.data[p.index:p.index+num]
  p.index += num
  return val
}

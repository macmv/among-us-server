package packet

import (
  "math"
  "encoding/binary"
)

type Packet struct {
  data []byte
  index int
}

func NewPacketFromBytes(bytes []byte) *Packet {
  p := Packet{}
  p.data = bytes
  return &p
}

func (p *Packet) Remaining() []byte {
  return p.data[p.index:]
}

func (p *Packet) ReadByte() byte {
  val := p.data[p.index]
  p.index += 1
  return val
}

func (p *Packet) ReadShort() uint16 {
  val := binary.LittleEndian.Uint16(p.data[p.index:p.index+2])
  p.index += 2
  return val
}

func (p *Packet) ReadInt() uint32 {
  val := binary.LittleEndian.Uint32(p.data[p.index:p.index+4])
  p.index += 4
  return val
}

func (p *Packet) ReadLong() uint64 {
  val := binary.LittleEndian.Uint64(p.data[p.index:p.index+8])
  p.index += 8
  return val
}

func (p *Packet) ReadFloat() float32 {
  val := math.Float32frombits(binary.LittleEndian.Uint32(p.data[p.index:p.index+4]))
  p.index += 4
  return val
}

func (p *Packet) ReadDouble() float64 {
  val := math.Float64frombits(binary.LittleEndian.Uint64(p.data[p.index:p.index+8]))
  p.index += 8
  return val
}

func (p *Packet) ReadString() string {
  length := int(p.data[p.index] + 1)
  val := string(p.data[p.index+1:p.index+length])
  p.index += length
  return val
}

func (p *Packet) ReadBytes(num int) []byte {
  val := p.data[p.index:p.index+num]
  p.index += num
  return val
}

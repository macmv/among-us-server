package connection

import (
  "fmt"
  "net"

  "github.com/macmv/among-us-server/packet"
)

type connection struct {
  conn *net.UDPConn
  addr *net.UDPAddr
}

func new_connection(conn *net.UDPConn, addr *net.UDPAddr) *connection {
  c := connection{}
  c.conn = conn
  c.addr = addr
  return &c
}

func (c *connection) handle(data []byte) bool {
  p := packet.NewIncomingPacketFromBytes(data)
  fmt.Println(data)
  switch p.Id() {
  case 8:
    p.ReadShort()
    p.ReadShort()
    p.ReadByte()
    p.ReadByte()
    p.ReadByte()
    name := p.ReadString()
    fmt.Println("name: ", name)
    fmt.Println("Got conneciton packet")
    out := packet.NewOutgoingPacket()
    out.WriteByte(0x00)
    out.WriteByte(0x52)
    out.WriteByte(0x00)
    out.WriteByte(0x0e)
    out.WriteByte(0x01)
    out.WriteByte(0x01) // num of servers
    out.WriteByte(0x11)
    out.WriteByte(0x00)
    out.WriteByte(0x00)
    out.WriteString("Master-4")
    out.WriteIP("198.58.115.57")
    out.WriteByte(0x07)
    out.WriteByte(0x56)
    out.WriteByte(0xc8) // ping?
    out.WriteByte(0x12)
    c.send_packet(out)

    out = packet.NewOutgoingPacket()
    out.WriteByte(0x0a)
    out.WriteByte(0x00)
    out.WriteByte(0x01)
    out.WriteByte(0xff)
    c.send_packet(out)
    // c.send_disconnect("MEMES")
    return false
  case 9:
    fmt.Println("Got disconnect")
    return true
  }
  fmt.Println("Unknown packet id:", p.Id())
  return false
}

func (c *connection) send_packet(out *packet.OugoingPacket) {
  out.Send(c.conn, c.addr)
}

func (c *connection) send_disconnect(reason string) {
  res := packet.NewOutgoingPacket()
  res.WriteByte(0x09)
  res.WriteByte(0x01)
  res.WriteByte(0x45)
  res.WriteByte(0x00)
  res.WriteByte(0x01)
  res.WriteByte(0x08)
  res.WriteString(reason)
  res.Send(c.conn, c.addr)
}

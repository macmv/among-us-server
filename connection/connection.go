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
    fmt.Println("name: ", p.ReadString())
    fmt.Println("Got conneciton packet")
    c.send_disconnect("MEMES")
    return false
  case 9:
    fmt.Println("Got disconnect")
    return true
  }
  fmt.Println("Unknown packet id:", p.Id())
  return false
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

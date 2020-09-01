package connection

import (
  "fmt"
  "net"

  "github.com/macmv/among-us-server/game"
  "github.com/macmv/among-us-server/packet"
)

type Connection struct {
  conn *net.UDPConn
  addr *net.UDPAddr
  player *game.Player
}

func new_connection(conn *net.UDPConn, addr *net.UDPAddr) *Connection {
  c := Connection{}
  c.conn = conn
  c.addr = addr
  return &c
}

func (c *Connection) handle(game *game.Game, data []byte) bool {
  p := packet.NewIncomingPacketFromBytes(data)
  fmt.Println("Got data:", data)
  switch p.Id() {
  case 8:
    p.ReadShort()
    p.ReadShort()
    p.ReadByte()
    p.ReadByte()
    p.ReadByte()
    name := p.ReadString()

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
    c.SendPacket(out)

    c.player = game.AddPlayer(name, c.conn, c.addr)
    return false
  case 9:
    return true
  }
  return false
}

func (c *Connection) SendPacket(out *packet.OutgoingPacket) {
  out.Send(c.conn, c.addr)
}

func (c *Connection) send_disconnect(reason string) {
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

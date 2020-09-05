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
  packet_type := p.ReadByte()
  if packet_type == 9 {
    return true
  } else if packet_type == 10 {
    // packet was recieved
    return false
  }
  p.ReadByte()
  id := p.ReadByte()

  out := packet.NewOutgoingPacket()
  out.WriteByte(0x0a)
  out.WriteByte(0x00)
  out.WriteByte(id)
  out.WriteByte(0xff)
  c.SendPacket(out)

  switch packet_type {
  case 1: // ingame packet
    val := p.ReadByte()
    if val == 5 { // wants to join a server
      p.ReadShort()
      c.player.JoinGame(p.ReadBytes(4))
    } else if val == 21 { // join game
      p.ReadByte()
      p.ReadByte()
      p.ReadInt()
      p.ReadShort()
      p.ReadBytes(4)
      game_type := p.ReadString()
      fmt.Println("Client is joining game with type ", game_type)
    } else if val == 44 { // listing all servers
      c.player.SendServerList()
    }
    return false
  case 8: // connect packet
    p.ReadByte()
    p.ReadByte()
    p.ReadByte()
    p.ReadByte()
    p.ReadByte()
    name := p.ReadString()
    fmt.Println("Name:", name)

    c.player = game.AddPlayer(name, c.conn, c.addr)
    return false
  case 12: // ping
    return false
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

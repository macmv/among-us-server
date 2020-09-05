package connection

import (
  "fmt"
  "net"

  "github.com/macmv/among-us-server/game"
  "github.com/macmv/among-us-server/packet"
  "github.com/macmv/among-us-server/packet_stream"
)

type Connection struct {
  player *game.Player
  outgoing_packets *packet_stream.OutgoingPacketStream
  // incoming_packets *packet_stream.IncomingPacketStream
}

func new_connection(conn *net.UDPConn, addr *net.UDPAddr) *Connection {
  c := Connection{}
  c.outgoing_packets = packet_stream.NewOutgoingPacketStream(conn, addr)
  // c.incoming_packets = packet_stream.NewIncomingPacketStream(conn, addr)
  return &c
}

func (c *Connection) handle(game *game.Game, data []byte) bool {
  p := packet.NewIncomingPacketFromBytes(data)
  packet_type := p.ReadByte()
  if packet_type == 9 {
    return true
  }
  p.ReadByte()
  id := p.ReadByte()

  // confirm packet
  if packet_type == 10 {
    c.outgoing_packets.Confirm(id)
    return false
  }

  fmt.Println("Got data:", data)

  out := packet.NewOutgoingPacket()
  out.WriteByte(0x0a)
  out.WriteByte(0x00)
  out.WriteByte(id)
  out.WriteByte(0xff)
  // DO NOT DO THIS
  // This is ok because the packet does not need to be confirmed.
  // Use c.SendPacket(out) for all other packets.
  c.outgoing_packets.SendNoConfirm(out)

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
      fmt.Println("Client is joining game with type", game_type)
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

    c.player = game.AddPlayer(name, c.outgoing_packets)
    return false
  case 12: // ping
    return false
  }
  return false
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

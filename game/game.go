package game

import (
  "net"

  "github.com/macmv/among-us-server/packet"
)

type Game struct {
  players map[string]*Player
}

func New() *Game {
  g := Game{}
  g.players = make(map[string]*Player)
  return &g
}

func (g *Game) AddPlayer(name string, conn *net.UDPConn, addr *net.UDPAddr) *Player {
  _, ok := g.players[name]
  if ok {
    return nil
  }
  p := new_player(name, conn, addr)
  g.players[name] = p

  out := packet.NewOutgoingPacket()
  out.WriteByte(0x0a)
  out.WriteByte(0x00)
  out.WriteByte(0x01)
  out.WriteByte(0xff)
  p.SendPacket(out)

  return p
}

func (g *Game) RemovePlayer(name string) {
  delete(g.players, name)
}

package game

import (
  "github.com/macmv/among-us-server/packet_stream"
)

type Game struct {
  players map[string]*Player
}

func New() *Game {
  g := Game{}
  g.players = make(map[string]*Player)
  go (&g).startUpdateLoop()
  return &g
}

func (g *Game) AddPlayer(name string, outgoing_packets *packet_stream.OutgoingPacketStream) *Player {
  _, ok := g.players[name]
  if ok {
    return nil
  }
  p := new_player(name, outgoing_packets)
  g.players[name] = p

  return p
}

func (g *Game) RemovePlayer(name string) {
  delete(g.players, name)
}

func (g *Game) startUpdateLoop() {
}

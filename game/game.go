package game

import (
  "net"
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

func (g *Game) AddPlayer(name string, conn *net.UDPConn, addr *net.UDPAddr) *Player {
  _, ok := g.players[name]
  if ok {
    return nil
  }
  p := new_player(name, conn, addr)
  g.players[name] = p

  return p
}

func (g *Game) RemovePlayer(name string) {
  delete(g.players, name)
}

func (g *Game) startUpdateLoop() {
}

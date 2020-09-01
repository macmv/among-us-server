package game

import (
  "net"

  "github.com/macmv/among-us-server/packet"
)

type Player struct {
  name string
  conn *net.UDPConn
  addr *net.UDPAddr
}

func new_player(name string, conn *net.UDPConn, addr *net.UDPAddr) *Player {
  p := Player{}
  p.name = name
  p.conn = conn
  p.addr = addr
  return &p
}

func (p *Player) Name() string {
  return p.name
}

func (p *Player) SendPacket(out *packet.OutgoingPacket) {
  out.Send(p.conn, p.addr)
}

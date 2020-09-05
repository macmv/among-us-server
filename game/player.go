package game

import (
  "github.com/macmv/among-us-server/packet"
  "github.com/macmv/among-us-server/packet_stream"
)

type Player struct {
  name string
  outgoing_packets *packet_stream.OutgoingPacketStream
}

func new_player(name string, outgoing_packets *packet_stream.OutgoingPacketStream) *Player {
  p := Player{}
  p.name = name
  p.outgoing_packets = outgoing_packets
  return &p
}

func (p *Player) Name() string {
  return p.name
}

func (p *Player) SendPacket(out *packet.OutgoingPacket) {
  out.Send(p.conn, p.addr)
}

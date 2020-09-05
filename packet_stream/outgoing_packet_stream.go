package packet_stream

import (
  "net"
  "fmt"
  "time"

  "github.com/macmv/among-us-server/packet"
)

type packet_time struct {
  packet *packet.OutgoingPacket
  time time.Time
}

type OutgoingPacketStream struct {
  outgoing_packets map[byte]packet_time
  outgoing_counter byte
  conn *net.UDPConn
  addr *net.UDPAddr
}

func NewOutgoingPacketStream(conn *net.UDPConn, addr *net.UDPAddr) *OutgoingPacketStream {
  o := OutgoingPacketStream{}
  o.outgoing_packets = make(map[byte]packet_time)
  return &o
}

func (o *OutgoingPacketStream) Send(out *packet.OutgoingPacket) {
  o.outgoing_counter++
  o.outgoing_packets[o.outgoing_counter] = packet_time{packet: out, time: time.Now()}
  // Check if any packets that are over a second old have not gotten confirm
  for id, packet := range o.outgoing_packets {
    if time.Since(packet.time) > time.Second {
      fmt.Println("Resending id", id, "data", out.Data())
      packet.packet.Send(o.conn, o.addr)
    }
  }
  fmt.Println("Sending data", out.Data())
  o.SendNoConfirm(out)
}

func (o *OutgoingPacketStream) SendNoConfirm(out *packet.OutgoingPacket) {
  out.Send(o.conn, o.addr)
}

func (o *OutgoingPacketStream) Confirm(id byte) {
  delete(o.outgoing_packets, id)
}

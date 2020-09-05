package connection

import (
  "net"

  pb "github.com/macmv/among-us-server/proto"
)

func (c *Connection) SendBroadcast(addr *net.UDPAddr, data []byte) {

}

func (c *Connection) SendGame(addr *net.UDPAddr, data []byte) {
  out := pb.Packet{}
  out.Addr = addr.String()
  out.Data = data
  c.stream.Send(&out)
}

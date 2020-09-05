package game

import (
  "fmt"

  "github.com/macmv/among-us-server/packet"
)

func (p *Player) JoinGame(code []byte) {
  fmt.Println("Adding player to game")
  out := packet.NewOutgoingPacket()
  out.WriteByte(0x01)
  out.WriteByte(0x00)
  out.WriteByte(0x01)
  out.WriteByte(0x22)
  out.WriteByte(0x00)
  out.WriteByte(0x07)
  out.WriteBytes(code)
  out.WriteByte(0x5a)
  out.WriteByte(0xb4)
  out.WriteByte(0x04)
  out.WriteByte(0x00)
  out.WriteByte(0x0f)
  out.WriteByte(0xb0)
  out.WriteByte(0x04)
  out.WriteByte(0x00)
  out.WriteByte(0x07)
  out.WriteByte(0x5a)
  out.WriteByte(0x5a)
  p.SendPacket(out)
}

func (p *Player) SendPong(id byte) {
  out := packet.NewOutgoingPacket()
  out.WriteByte(0x0a)
  out.WriteByte(0x00)
  out.WriteByte(id)
  out.WriteByte(0xff)
  p.SendPacket(out)
}

func (p *Player) SendServerList() {
  fmt.Println("Sending player server list")
  out := packet.NewOutgoingPacket()
  out.WriteByte(0x01)
  out.WriteByte(0x00)
  out.WriteByte(0x02)
  out.WriteByte(0x80)

  out.WriteByte(0x04)
  out.WriteByte(0x10)
  out.WriteByte(0x7e)
  out.WriteByte(0x04)
  out.WriteByte(0x00)

  out.WriteBytes([]byte("AAAA"))
  out.WriteString("The server")

  // if i == 0 {
  //   fmt.Println("??                   :", packet.ReadByte())
  //   fmt.Println("??                   :", packet.ReadBytes(6))
  // }

  out.WriteByte(5)       // num players
  out.WriteByte(123)     // ping
  out.WriteByte(2)       // bitfield?
  out.WriteByte(10)      // max players
  out.WriteByte(2)       // impostors
  out.WriteDouble(1.125) // kill dist
  out.WriteFloat(1)      // crewmate vision
  out.WriteFloat(1.5)    // impostor vision
  out.WriteFloat(14)     // emergency cooldown
  out.WriteByte(1)
  out.WriteByte(1)
  out.WriteByte(1)
  out.WriteInt(1)
  out.WriteByte(1)
  out.WriteByte(1)
  out.WriteInt(30)       // voting time
  out.WriteInt(120)      // discussion time

  p.SendPacket(out)
}

func (p *Player) SendDisconnect(reason string) {

}

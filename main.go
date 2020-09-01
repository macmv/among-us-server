package main

import (
  "fmt"
  "net"
  "io/ioutil"
  "encoding/hex"

  "github.com/macmv/among-us-server/packet"
  "github.com/macmv/among-us-server/connection"
)

func main() {
  connection.Listen(":50000")
}

func handle_client(conn *net.UDPConn) {
  for {
    data := make([]byte, 1024)
    length, addr, err := conn.ReadFromUDP(data)
    fmt.Println(addr)
    if err != nil {
      panic(err)
    }
    fmt.Println(data[:length])
    handle_packet(data[:length], conn, addr)
  }
}

func handle_packet(data []byte, conn *net.UDPConn, addr *net.UDPAddr) {
  id := data[0]
  if id == 8 {
  }
}

func parse_server_info() {
  packet_str, _ := ioutil.ReadFile("packet-hex.txt")
  packet_bytes, _ := hex.DecodeString(string(packet_str))

  packet := packet.NewIncomingPacketFromBytes(packet_bytes)
  fmt.Println()
  i := 0
  for len(packet.Remaining()) > 0 {
    fmt.Println("I:                                   ", i)
    is_next := packet.ReadShort()
    fmt.Println("?? (short)           :", is_next)
    if i >= 20 {
      break
    }
    if i == 0 {
      fmt.Println("??                   :", packet.ReadByte())
      fmt.Println("??                   :", packet.ReadBytes(6))
    }
    fmt.Println("server code          :", string(packet.ReadBytes(4)))
    fmt.Println("server name          :", packet.ReadString())
    fmt.Println("num players          :", packet.ReadByte())
    fmt.Println("(ping?)              :", packet.ReadByte())
    flags := packet.ReadByte()
    fmt.Println("?? (bitfield?)       :", flags)
    val := packet.ReadByte()
    if val < 4 {
      fmt.Println("?? (not max players) :", val)
      fmt.Println("max players          :", packet.ReadByte())
    } else {
      fmt.Println("max players          :", val)
    }
    fmt.Println("num impostors        :", packet.ReadByte())
    fmt.Println("?? (kill dist?)      :", packet.ReadDouble())
    fmt.Println("crewmate vision      :", packet.ReadFloat())
    fmt.Println("impostor vision      :", packet.ReadFloat())
    fmt.Println("emergency cooldown   :", packet.ReadFloat())
    fmt.Println("(tasks?)             :", packet.ReadByte())
    fmt.Println("(tasks?)             :", packet.ReadByte())
    fmt.Println("(tasks?)             :", packet.ReadByte())
    fmt.Println("?? (int32)           :", packet.ReadInt())
    fmt.Println("??                   :", packet.ReadByte())
    fmt.Println("??                   :", packet.ReadByte())
    fmt.Println("discussion time      :", packet.ReadInt())
    fmt.Println("voting time          :", packet.ReadInt())
    i++
  }

  fmt.Println(packet.Remaining())
}

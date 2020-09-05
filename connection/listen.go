package connection

import (
  "net"
  "time"
)

func (c *Connection) StartBroadcast(local_port, port string) {
  message := []byte{}
  message = append(message, []byte{0x04, 0x02}...)
  message = append(message, []byte("My Server")...)
  message = append(message, []byte("~")...)
  message = append(message, []byte("Open")...)
  message = append(message, []byte("~")...)
  message = append(message, []byte("-1")...)
  message = append(message, []byte("~")...)

  local, err := net.ResolveUDPAddr("udp", local_port)
  if err != nil {
    panic(err)
  }
  remote, err := net.ResolveUDPAddr("udp", "255.255.255.255" + port)
  if err != nil {
    panic(err)
  }

  ticker := time.NewTicker(time.Second)
  for range ticker.C {
    conn, err := net.DialUDP("udp", local, remote)
    if err != nil {
      panic(err)
    }
    conn.Write(message)
    conn.Close()
  }
}

func (c *Connection) ListenGame(port string) {
  addr, err := net.ResolveUDPAddr("udp", port)
  if err != nil {
    panic(err)
  }
  ln, err := net.ListenUDP("udp", addr)
  if err != nil {
    panic(err)
  }
  defer ln.Close()

  buf := make([]byte, 4096)
  for {
    n, addr, err := ln.ReadFromUDP(buf)
    if err != nil {
      panic(err)
    }

    c.SendGame(addr, buf[:n])
  }
}

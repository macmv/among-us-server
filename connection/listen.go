package connection

import (
  "fmt"
  "net"
  "time"
)

func (c *Connection) OpenUDP(port string) *net.UDPConn {
  addr, err := net.ResolveUDPAddr("udp", port)
  if err != nil {
    panic(err)
  }
  ln, err := net.ListenUDP("udp", addr)
  if err != nil {
    panic(err)
  }
  return ln
}

func (c *Connection) StartBroadcast(conn *net.UDPConn, port string) {
  message := []byte{}
  message = append(message, []byte{0x04, 0x02}...)
  message = append(message, []byte("My Server")...)
  message = append(message, []byte("~")...)
  message = append(message, []byte("Open")...)
  message = append(message, []byte("~")...)
  message = append(message, []byte("-1")...)
  message = append(message, []byte("~")...)

  addr, err := net.ResolveUDPAddr("udp", "255.255.255.255" + port)
  if err != nil {
    panic(err)
  }

  ticker := time.NewTicker(time.Second)
  for range ticker.C {
    conn.WriteToUDP(message, addr)
  }
}

func (c *Connection) StartHost(addr *net.UDPAddr) {
  for {
    time.Sleep(time.Second)
  }
}

func (c *Connection) ListenBroadcast(conn *net.UDPConn) *net.UDPAddr {
  buf := make([]byte, 4096)
  for {
    _, addr, err := conn.ReadFromUDP(buf)
    if err != nil {
      panic(err)
    }
    return addr
  }
}

func (c *Connection) ListenGame(ln *net.UDPConn) {
  buf := make([]byte, 4096)
  for {
    n, addr, err := ln.ReadFromUDP(buf)
    if err != nil {
      panic(err)
    }

    fmt.Println("Got game packet", buf[:n])

    c.SendGame(addr, buf[:n])
  }
}

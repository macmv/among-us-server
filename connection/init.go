package connection

import (
  "fmt"
  "net"
)

func Listen(port string) {
  addr, err := net.ResolveUDPAddr("udp", port)
  if err != nil {
    panic(err)
  }
  fmt.Println("Listening on :50000")
  conn, err := net.ListenUDP("udp", addr)
  if err != nil {
    panic(err)
  }
  connections := make(map[*net.UDPAddr]*connection)
  for {
    data := make([]byte, 1024)
    length, addr, err := conn.ReadFromUDP(data)
    if err != nil {
      panic(err)
    }
    if connections[addr] == nil {
      connections[addr] = new_connection(conn, addr)
    }
    if connections[addr].handle(data[:length]) {
      fmt.Println("Closing connection with: ", addr)
      delete(connections, addr)
    }
  }
}

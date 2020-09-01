package connection

import (
  "fmt"
  "net"

  "github.com/macmv/among-us-server/game"
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
  connections := make(map[string]*Connection)
  g := game.New()
  for {
    data := make([]byte, 1024)
    length, addr, err := conn.ReadFromUDP(data)
    if err != nil {
      panic(err)
    }
    addr_string := addr.String()
    if connections[addr_string] == nil {
      connections[addr_string] = new_connection(conn, addr)
    }
    if connections[addr_string].handle(g, data[:length]) {
      g.RemovePlayer(connections[addr_string].player.Name())
      fmt.Println("Closing connection with: ", addr)
      delete(connections, addr_string)
    }
  }
}

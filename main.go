package main

import (
  "fmt"
  "net"
  "flag"
  "io/ioutil"
  "encoding/hex"

  "google.golang.org/grpc"
  pb "github.com/macmv/among-us-server/proto"
  "github.com/macmv/among-us-server/packet"
  "github.com/macmv/among-us-server/connection"
)

var (
  flag_server = flag.Bool("server", false, "Set if this should be the server")
  flag_host = flag.Bool("host", false, "Set if the user is going to host the game")
  flag_port = flag.String("port", ":7444", "Used to change thr grpc port")
)

func main() {
  flag.Parse()
  if *flag_server {
    start_server(*flag_port)
  } else {
    if *flag_host {
      start_host(*flag_port)
    } else {
      start_client(*flag_port)
    }
  }
}

func start_host(server_port string) {
  c := connection.New(server_port)
  for {
    fmt.Println("Listening for broadcasts on localhost...")
    conn := c.OpenUDP(":47777")
    addr := c.ListenBroadcast(conn)
    fmt.Println("Found game, listening for game packets")
    c.StartHost(addr)
    fmt.Println("Game closed, restarting")
  }
}

func start_client(server_port string) {
  c := connection.New(server_port)
  for {
    fmt.Println("Opening port 12345")
    conn := c.OpenUDP(":12345")
    fmt.Println("Broadcasting local game, and listening for game packets")
    go c.StartBroadcast(conn, ":47777")
    conn = c.OpenUDP(":22023")
    c.ListenGame(conn)
  }
}

func start_server(port string) {
  ln, err := net.Listen("tcp", port)
  if err != nil {
    panic(err)
  }
  server := grpc.NewServer()
  pb.RegisterAmongUsServer(server, new_server())
  server.Serve(ln)
}

type server struct {
  host_addr string
  connections map[pb.AmongUs_ConnectionServer]struct{}
}

func new_server() *server {
  s := server{}
  s.connections = make(map[pb.AmongUs_ConnectionServer]struct{})
  return &s
}

func (s *server) Connection(conn pb.AmongUs_ConnectionServer) error {
  s.connections[conn] = struct{}{}
  for {
    p, err := conn.Recv()
    if err != nil {
      panic(err)
    }
    fmt.Println("Got packet", p)
    for other_conn := range s.connections {
      if other_conn != conn {
        other_conn.Send(p)
      }
    }
  }
  return nil
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

package connection

import (
  "context"

  "google.golang.org/grpc"
  pb "github.com/macmv/among-us-server/proto"
)

func New(ip string) *Connection {
  conn, err := grpc.Dial(ip, grpc.WithInsecure())
  if err != nil {
    panic(err)
  }
  client := pb.NewAmongUsClient(conn)
  c := Connection{}
  c.stream, err = client.Connection(context.Background())
  if err != nil {
    panic(err)
  }
  return &c
}

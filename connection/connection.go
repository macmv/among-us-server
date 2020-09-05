package connection

import (
  pb "github.com/macmv/among-us-server/proto"
)

type Connection struct {
  stream pb.AmongUs_ConnectionClient
}


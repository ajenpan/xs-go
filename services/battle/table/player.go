package table

import (
	protobuf "google.golang.org/protobuf/proto"

	"xs/services/battle"
	pb "xs/services/battle/proto"
)

type player struct {
	*pb.PlayerInfo

	desk *Table
}

func NewPlayer() battle.Player {
	return &player{}
}

func (p *player) GetScore() int64 {
	return int64(p.Score)
}

func (p *player) GetUserID() int64 {
	return p.Uid
}

func (p *player) GetSeatID() battle.SeatID {
	return battle.SeatID(p.SeatId)
}

func (p *player) IsRobot() bool {
	return false
}

func (p *player) SendMessage(protobuf.Message) error {
	return nil
}

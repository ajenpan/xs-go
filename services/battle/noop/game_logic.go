package noop

import (
	"time"

	"google.golang.org/protobuf/proto"

	"xs/services/battle"
)

func NewGameLogic() battle.GameLogic {
	return &GameLogic{}
}

type GameLogic struct {
}

func (gl *GameLogic) OnInit(desk battle.GameDesk, conf interface{}) error { return nil }
func (gl *GameLogic) OnStart(players []battle.Player) error               { return nil }
func (gl *GameLogic) OnTick(time.Duration)                                {}
func (gl *GameLogic) OnReset()                                            {}
func (gl *GameLogic) OnMessage(battle.Player, string, []byte)             {}
func (gl *GameLogic) OnEvent(string, proto.Message)                       {}

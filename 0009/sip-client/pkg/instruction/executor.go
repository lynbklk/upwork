package instruction

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
)

type Executor struct {
	ctx context.Context
}

func NewExecutor(ctx context.Context) *Executor {
	return &Executor{
		ctx: ctx,
	}
}

func (e *Executor) Run() {
	for {
		ins := <-Chan
		switch ins.ActionType {
		case Dial:
			e.dial(ins.Who, ins.Address)
		case Wait:
			e.wait(ins.Who, ins.Address, ins.Any)
		case PlayFromFile:
			e.play(ins.File)
		case RecordToFile:
			e.record(ins.File)
		case DTMF:
			e.dtmf(ins.Num)
		case Exit:
			e.exit()
		default:
			log.Info().Msgf("ignore unknown instruction, type: %v", ins.ActionType)
			continue
		}
	}
}

func (e *Executor) dial(who string, address string) error {
	return nil
}

func (e *Executor) wait(who string, address string, any bool) error {
	return nil
}

func (e *Executor) play(file string) error {
	return nil
}

func (e *Executor) record(file string) error {
	return nil
}

func (e *Executor) dtmf(num int64) error {
	return nil
}

func (e *Executor) exit() {
	log.Info().Msg("exit right now.")
	os.Exit(0)
}

package instruction

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
)

var (
	Chan chan *Instruction
)

type Reader struct {
	ctx context.Context
}

func NewReader(ctx context.Context) *Reader {
	return &Reader{
		ctx: ctx,
	}
}

func (r *Reader) Run() {
	var input string

	for {
		fmt.Scan(&input)
		log.Info().Msgf("instruction received: %s", input)
		ins, err := r.resolve(input)
		if err != nil {
			log.Info().Err(err).Msgf("ignore illegal command: %s", input)
			continue
		}
		Chan <- ins
	}
}

func (r *Reader) resolve(input string) (*Instruction, error) {
	instruct := &Instruction{}
	if input == "exit" {
		instruct.ActionType = Exit
		return instruct, nil
	}
	num, err := strconv.ParseInt(input, 10, 64)
	if err == nil {
		instruct.ActionType = DTMF
		instruct.Num = num
		return instruct, nil
	}
	err = fmt.Errorf("illegal instruction")
	if strings.HasPrefix(strings.ToLower(input), "dial:") {
		splits := strings.Split(input, ":")
		if len(splits) != 2 {
			instruct.ActionType = Illegal
			return instruct, err
		}
		splits = strings.Split(splits[1], "@")
		if len(splits) != 2 {
			instruct.ActionType = Illegal
			return instruct, err
		}
		instruct.ActionType = Dial
		instruct.Who = splits[0]
		instruct.Address = splits[1]
		return instruct, nil
	}
	if strings.HasPrefix(strings.ToLower(input), "wait") {
		splits := strings.Split(input, ":")
		if len(splits) != 2 {
			instruct.ActionType = Illegal
			return instruct, err
		}
		splits = strings.Split(splits[1], "@")
		length := len(splits)
		if length == 1 {
			if splits[0] == "any" {
				instruct.ActionType = Wait
				instruct.Any = true
			} else {
				instruct.ActionType = Illegal
			}
			return instruct, nil
		}
		if length == 2 {
			instruct.ActionType = Wait
			instruct.Who = splits[0]
			instruct.Address = splits[1]
			return instruct, nil
		}
		instruct.ActionType = Illegal
		return instruct, err
	}
	if strings.HasPrefix(strings.ToLower(input), "playfromfile") {
		splits := strings.Split(input, ":")
		if len(splits) != 2 {
			instruct.ActionType = Illegal
			return instruct, err
		}
		instruct.ActionType = PlayFromFile
		instruct.File = splits[1]
		return instruct, nil
	}
	if strings.HasPrefix(strings.ToLower(input), "recordtofile") {
		splits := strings.Split(input, ":")
		if len(splits) != 2 {
			instruct.ActionType = Illegal
			return instruct, err
		}
		instruct.ActionType = RecordToFile
		instruct.File = splits[1]
		return instruct, nil
	}
	instruct.ActionType = Illegal
	return instruct, err
}

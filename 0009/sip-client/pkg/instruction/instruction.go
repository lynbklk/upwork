package instruction

const (
	Dial         ActionType = 1
	Wait         ActionType = 2
	PlayFromFile ActionType = 3
	RecordToFile ActionType = 4
	DTMF         ActionType = 5
	Exit         ActionType = 6
	Illegal      ActionType = 7
)

type ActionType int

type Instruction struct {
	ActionType ActionType
	Who        string
	Address    string
	File       string
	Num        int64
	Any        bool
}

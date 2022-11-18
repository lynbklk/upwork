package timer

import (
	"flag"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var (
	second int64
)

func Start(second int64) {
	if second == 0 {
		return
	}
	log.Info().Msgf("sip will exit automate after %d seconds.", second)
	timer := time.NewTimer(time.Duration(second))
	select {
	case <-timer.C:
		log.Info().Msgf("sip exit for timeout.")
		os.Exit(0)
	}
}

func init() {
	flag.Int64Var(&second, "seconds", 0, "sip seconds. 0 means no limit.")
	go Start(second)
}

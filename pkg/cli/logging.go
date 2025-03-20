package cli

import (
	"encoding/json"
	stdlog "log"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
)

func InitLogging(level zerolog.Level) {
	log.Logger = log.Level(level).Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
		FormatTimestamp: func(i any) string {
			num, err := oops.Wrap2(i.(json.Number).Int64())
			if err != nil {
				stdlog.Fatalf("%+v", err)
			}
			timestamp := time.Unix(num, 0)
			return timestamp.Format(time.RFC3339)
		},
	})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

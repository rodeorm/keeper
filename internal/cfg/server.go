package cfg

import "github.com/rodeorm/keeper/internal/sender"

func ConfigurateServer(configFile string, q *sender.Queue) (Server, error) {
	builder := &ServerBuilder{}

	return builder.SetConfig(configFile).
		Build(), nil
}

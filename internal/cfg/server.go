package cfg

import "github.com/rodeorm/keeper/internal/core"

func ConfigurateServer(configFile string) (*Server, error) {
	queue := core.NewQueue(3)

	builder := &ServerBuilder{}
	srv := builder.SetConfig(configFile).
		Build()
	srv.MessageQueue = queue

	return &srv, nil
}

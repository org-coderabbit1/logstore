package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/prometheus/common/version"

	"example.com/acme/logstore/v3/pkg/tool/commands"
)

var (
	ruleCommand  commands.RuleCommand
	auditCommand commands.AuditCommand
)

func main() {
	app := kingpin.New("logstoretool", "A command-line tool to manage Logstore.")
	ruleCommand.Register(app)
	auditCommand.Register(app)

	app.Command("version", "Get the version of the logstoretool CLI").Action(func(_ *kingpin.ParseContext) error {
		fmt.Println(version.Print("logstore"))
		return nil
	})

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

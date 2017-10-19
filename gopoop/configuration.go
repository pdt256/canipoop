package gopoop

import (
	"github.com/namsral/flag"
	"os"
	"strings"
)

type Configuration struct {
	storageBucket string
	rooms         string
}

func GetFlagConfiguration() (config Configuration) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	flag.String(
		flag.DefaultConfigFlagname,
		``,
		`path to config file`,
	)
	flag.StringVar(
		&config.storageBucket,
		`storageBucket`,
		`canipoop-4efd0`,
		`firebase storage bucket`,
	)
	flag.StringVar(
		&config.rooms,
		`rooms`,
		`office1/br2,office1/br1`,
		`room(s) location`,
	)
	flag.Parse()

	return
}

func (configuration Configuration) GetRooms() []string {
	return strings.Split(configuration.rooms, `,`)
}

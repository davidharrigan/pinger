// +build integration

package pinger

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/davidharrigan/pinger/testing/config"
)

var integrationLocal = flag.Bool("integration-local", false, "spin up a local instance of pinger within the test")
var pingerAddress = flag.String("pinger-address", "localhost", "pinger address")
var pingerPort = flag.Int("pinger-port", 50051, "pinger port")

func TestMain(m *testing.M) {
	flag.Parse()

	if *integrationLocal {
		go func() {
			s := config.Pinger()
			defer s.Stop()
		}()
	} else {
		log.Println(">>> running in integration mode")
		config.PingerConfig = config.ServerConfig{
			Address: *pingerAddress,
			Port:    *pingerPort,
		}
	}

	os.Exit(m.Run())
}

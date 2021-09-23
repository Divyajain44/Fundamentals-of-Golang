package test

import (
	"os"
	"testing"
	"testing/golang-testing/src/api/app"

	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}

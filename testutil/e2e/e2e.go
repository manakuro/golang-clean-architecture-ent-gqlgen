package e2e

import (
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/pkg/adapter/controller"
	"golang-clean-architecture-ent-gqlgen/pkg/infrastructure/graphql"
	"golang-clean-architecture-ent-gqlgen/pkg/infrastructure/router"
	"golang-clean-architecture-ent-gqlgen/pkg/registry"
	"golang-clean-architecture-ent-gqlgen/testutil"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

// SetupOption is an option of SetupE2E
type SetupOption struct {
	Teardown func(t *testing.T, client *ent.Client)
}

// Setup set up database and server for E2E test
func Setup(t *testing.T, option SetupOption) (expect *httpexpect.Expect, client *ent.Client, teardown func()) {
	t.Helper()
	testutil.ReadConfigE2E()

	client = testutil.NewDBClient(t)
	ctrl := newController(client)
	gqlsrv := graphql.NewServer(client, ctrl)
	e := router.New(gqlsrv)

	srv := httptest.NewServer(e)

	return httpexpect.WithConfig(httpexpect.Config{
			BaseURL:  srv.URL,
			Reporter: httpexpect.NewAssertReporter(t),
			Printers: []httpexpect.Printer{
				httpexpect.NewDebugPrinter(t, true),
			},
		}), client, func() {
			option.Teardown(t, client)
			defer client.Close()
			defer srv.Close()
		}
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}

// GetData gets data from graphql response.
func GetData(e *httpexpect.Response) *httpexpect.Value {
	return e.JSON().Path("$.data")
}

// GetObject return data from path.
// Path returns a new Value object for child object(s) matching given
// JSONPath expression.
// Example 1:
//  json := `{"users": [{"name": "john"}, {"name": "bob"}]}`
//  value := NewValue(t, json)
//
//  value.Path("$.users[0].name").String().Equal("john")
//  value.Path("$.users[1].name").String().Equal("bob")
func GetObject(obj *httpexpect.Object, path string) *httpexpect.Object {
	return obj.Path("$." + path).Object()
}

// GetErrors return errors from graphql response.
func GetErrors(e *httpexpect.Response) *httpexpect.Value {
	return e.JSON().Path("$.errors")
}

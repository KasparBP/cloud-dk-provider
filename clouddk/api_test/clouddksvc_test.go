package api_test

import (
	"context"
	"github.com/KasparBP/cloud-dk-provider/clouddk/api"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func loadTestData(t *testing.T, fname string) []byte {
	path := filepath.Join("testdata", fname)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

// Setup test server which will respond with testdata loaded from supplied filename
func setupTestServer(t *testing.T, fname string) *httptest.Server {
	response := loadTestData(t, fname)
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		_, err := rw.Write(response)
		if err != nil {
			t.Fatal(err)
		}
	}))
}

func testClient(t *testing.T, s *httptest.Server) *api.Client  {
	client, err := api.NewClient("token", s.URL)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestGetCloudServer(t *testing.T) {
	server := setupTestServer(t, "getcloudserver_resp.json")
	defer server.Close()

	client := testClient(t, server)
	response, err := client.ClouddkService.GetCloudServer(context.Background(), "identifier1")
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, bool(response.Booted))
	assert.Equal(t, response.Identifier, "identifier1")
}

func TestCreateCloudServer(t *testing.T) {
	server := setupTestServer(t, "createcloudserver_resp.json")
	defer server.Close()

	client := testClient(t, server)
	cs, err := client.ClouddkService.CreateCloudServer(context.Background(), &api.CloudServer{
		HostName:            "hostname",
		Label:               "label",
		InitialRootPassword: nil,
		Template: api.Template{
			Identifier: "ubuntu-18.04-x64",
		},
		Location: api.Location{
			Identifier: "dk1",
		},
		Package: api.Package{
			Identifier: "ac949a1cb4731d",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, cs.Identifier, "identifier1")
}

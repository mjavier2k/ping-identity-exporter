package ping_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/mjavier2k/ping-identity-exporter/pkg/ping"
	"gopkg.in/h2non/gock.v1"
)

var (
	pingHostname = "https://lv1testpingaccess01.example.com:3000/pa/heartbeat.ping"
	pingClient   = ping.Client{
		Hostname:   pingHostname,
		Endpoint:   fmt.Sprintf("%v", pingHostname),
		HTTPClient: &http.Client{},
	}
)

func TestClient_GetPingAccessHearthbeat(t *testing.T) {
	fixture, err := ioutil.ReadFile("../../test/fixtures/heartbeat.json")
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		s       ping.Client
		want    string
		wantErr bool
	}{
		{
			name: "GetPingAccessHeartbeat response should match fixture",
			want: "0.69",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			gock.Observe(gock.DumpRequest)
			gock.New(pingClient.Hostname).
				MatchType("json").
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := pingClient.GetPingAccessHearthbeat()
			fmt.Println(gotRaw)
			got := gotRaw.Items[0].CPULoad

			if (err != nil) != tt.wantErr {
				t.Errorf("GetPingAccessHearthbeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPingAccessHearthbeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

package ping

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

func NewPingClient(hostname string) *Client {

	insecure := viper.GetBool(InsecureSSLFlag)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	return &Client{
		Hostname: hostname,
		HTTPClient: &http.Client{
			Transport: tr,
			Timeout:   time.Duration(viper.GetInt64(HTTPClientTimeoutFlag)) * time.Second,
		},
		PingAccessHeartbeatEndpoint: fmt.Sprintf(viper.GetString(PingAccessHeartbeatEndpoint), hostname),
	}
}

func doHTTPGet(c *Client, body []byte) ([]byte, error) {
	log.Infof("Start scrape for %v\n", c.Hostname)
	req, err := http.NewRequest("GET", c.PingAccessHeartbeatEndpoint, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making RPC call %v: %v", string(body), err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Received invalid status code from RPC call: %v", resp.StatusCode)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	log.Infof("Completed scrape on %v", c.Hostname)
	return body, nil
}

// GetPingAccessHearthbeat is
func (c *Client) GetPingAccessHearthbeat() (PingAccessHBResponse, error) {
	r := PingAccessHBResponse{}
	bodyBytes, err := doHTTPGet(c, nil)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

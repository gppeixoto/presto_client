package presto_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const (
	healthQuery = "SHOW TABLES"
	maxRetries  = 5
)

type httpPrestoClient struct {
	client *http.Client
	addr   string
	logger *zap.Logger
}

// NewHttpPrestoClient returns a Presto HTTP client.
func NewHttpPrestoClient(addr string, logger *zap.Logger) (PrestoClient, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := httpPrestoClient{
		addr:   addr,
		logger: logger,
		client: &http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}

	return &client, nil
}

func (c *httpPrestoClient) WorkersCount() (int, error) {
	c.logger.Info("checking WorkersCount")
	var errorMsg string

	resp, err := c.client.Get(c.addr + "/v1/node")
	if err != nil {
		errorMsg = "unable to create GET for node route: " + err.Error()
		c.logger.Error(errorMsg)
		return 0, errors.New(errorMsg)
	}
	defer resp.Body.Close()

	nodeResp := nodeResponse{}
	if err := unmarshalJson(resp, nodeResp); err != nil {
		errorMsg = "unable to parse json " + err.Error()
		c.logger.Error(errorMsg)
		return 0, errors.New(errorMsg)
	}

	return len(nodeResp), nil
}

func (c *httpPrestoClient) Healthcheck() error {
	postForPresto, err := mkRequest(c.addr)
	if err != nil {
		return err
	}

	res, err := c.client.Do(postForPresto)
	switch {
	case err != nil:
		return err
	case res.StatusCode != 200:
		return errors.New("presto did not respond with 200 " + err.Error())
	}

	r := &apiResponse{}
	if err = unmarshalJson(res, r); err != nil {
		return err
	}

	return healthcheck(c, r)
}

func healthcheck(c *httpPrestoClient, r *apiResponse) error {
	res, err := c.client.Get(r.NextURI)
	if err != nil {
		return err
	}

	qr := &queryResult{}
	if err = unmarshalJson(res, qr); err != nil {
		return err
	}

	try := 0
	for ; try < maxRetries || len(qr.Columns) == 0; try++ {
		res, _ := c.client.Get(qr.NextURI)
		unmarshalJson(res, qr)
	}

	if try < maxRetries || len(qr.Columns) > 0 {
		return nil
	}

	return errors.New("exceeded retry limit, Presto might be down")
}

func unmarshalJson(res *http.Response, i interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(&i)
}

func mkRequest(addr string) (*http.Request, error) {
	request, err := http.NewRequest("POST",
		addr+"/v1/statement",
		bytes.NewBufferString(healthQuery),
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set("x-presto-user", "warehouse-healthchecker")
	request.Header.Set("content-type", "text/plain")
	request.Header.Set("x-presto-catalog", "hive")
	request.Header.Set("x-presto-schema", "default")

	return request, nil
}

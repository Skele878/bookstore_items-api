package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/Skele878/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

// elasticsearch client nterface
type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}
type esClient struct {
	client *elastic.Client
}

// Obtain a client for an Elasticsearch cluster of two nodes,
// running on 127.0.0.1 (as a localhost in this case i change addr). Do not run the sniffer.
// Set the healthcheck interval to 10s. When requests fail,
// retry 5 times. Print error messages to os.Stderr and informational
// messages to os.Stdout.  [olivere github -> elastic -> wiki -> configuretion]

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		// elastic.SetHeaders(http.Header{
		// 	"X-Caller-Id": []string{"..."},
		// }),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

// setclient function
func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

// index function
func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index("items").
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

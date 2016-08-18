package controllers_test

import (
	"io"
	"net/http"

	"github.com/RackHD/ipam/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"

	"testing"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

func NewRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	Expect(err).NotTo(HaveOccurred())
	return req
}

func Do(req *http.Request) *http.Response {
	client := &http.Client{}

	res, err := client.Do(req)
	Expect(err).NotTo(HaveOccurred())

	return res
}

type MockIpam struct {
	Err     error
	Pools   []models.Pool
	Created models.Pool
	Updated models.Pool
	Deleted string
}

func NewMockIpam() *MockIpam {
	return &MockIpam{
		Err: nil,
		Pools: []models.Pool{
			{
				ID:   bson.ObjectIdHex("578af30bbc63780007d99195"),
				Name: "Mock Pool",
				Tags: []string{"mock"},
			},
		},
	}
}

// GetPools ...
func (mock *MockIpam) GetPools() ([]models.Pool, error) {
	if mock.Err != nil {
		return []models.Pool{}, mock.Err
	}

	return mock.Pools, mock.Err
}

// GetPool ...
func (mock *MockIpam) GetPool(id string) (models.Pool, error) {
	if mock.Err != nil {
		return models.Pool{}, mock.Err
	}

	return mock.Pools[0], mock.Err
}

// CreatePool ...
func (mock *MockIpam) CreatePool(pool models.Pool) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.Created = pool

	return mock.Err
}

// UpdatePool ...
func (mock *MockIpam) UpdatePool(pool models.Pool) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.Updated = pool

	return mock.Err
}

// DeletePool ...
func (mock *MockIpam) DeletePool(id string) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.Deleted = id

	return mock.Err
}

// GetSubnets ...
func (mock *MockIpam) GetSubnets() ([]models.Subnet, error) {
	return []models.Subnet{}, nil
}

// GetSubnet ...
func (mock *MockIpam) GetSubnet(string) (models.Subnet, error) {
	return models.Subnet{}, nil
}

// CreateSubnet ...
func (mock *MockIpam) CreateSubnet(models.Subnet) error {
	return nil
}

// UpdateSubnet ...
func (mock *MockIpam) UpdateSubnet(models.Subnet) error {
	return nil
}

// DeleteSubnet ...
func (mock *MockIpam) DeleteSubnet(string) error {
	return nil
}

// GetPoolSubnets ...
func (mock *MockIpam) GetPoolSubnets(string) ([]models.Subnet, error) {
	return []models.Subnet{}, nil
}

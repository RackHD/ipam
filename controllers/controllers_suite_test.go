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
	Err error

	Pools       []models.Pool
	PoolCreated models.Pool
	PoolUpdated models.Pool
	PoolDeleted string

	Subnets       []models.Subnet
	SubnetCreated models.Subnet
	SubnetUpdated models.Subnet
	SubnetDeleted string

	Reservations       []models.Reservation
	ReservationCreated models.Reservation
	ReservationUpdated models.Reservation
	ReservationDeleted string

	Leases       []models.Lease
	LeaseUpdated models.Lease
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
		Subnets: []models.Subnet{
			{
				ID:   bson.ObjectIdHex("578af30bbc63780007d99195"),
				Name: "Mock Subnet",
				Tags: []string{"mock"},
				Pool: bson.ObjectIdHex("578af30bbc63780007d99195"),
			},
		},
		Reservations: []models.Reservation{
			{
				ID:     bson.ObjectIdHex("578af30bbc63780007d99195"),
				Name:   "Mock Subnet",
				Tags:   []string{"mock"},
				Subnet: bson.ObjectIdHex("578af30bbc63780007d99195"),
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

	mock.PoolCreated = pool

	return mock.Err
}

// UpdatePool ...
func (mock *MockIpam) UpdatePool(pool models.Pool) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.PoolUpdated = pool

	return mock.Err
}

// DeletePool ...
func (mock *MockIpam) DeletePool(id string) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.PoolDeleted = id

	return mock.Err
}

// GetSubnets ...
func (mock *MockIpam) GetSubnets(string) ([]models.Subnet, error) {
	if mock.Err != nil {
		return []models.Subnet{}, mock.Err
	}

	return mock.Subnets, mock.Err
}

// GetSubnet ...
func (mock *MockIpam) GetSubnet(string) (models.Subnet, error) {
	if mock.Err != nil {
		return models.Subnet{}, mock.Err
	}

	return mock.Subnets[0], mock.Err
}

// CreateSubnet ...
func (mock *MockIpam) CreateSubnet(subnet models.Subnet) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.SubnetCreated = subnet

	return mock.Err
}

// UpdateSubnet ...
func (mock *MockIpam) UpdateSubnet(subnet models.Subnet) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.SubnetUpdated = subnet

	return mock.Err
}

// DeleteSubnet ...
func (mock *MockIpam) DeleteSubnet(id string) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.SubnetDeleted = id

	return mock.Err
}

// GetReservations ...
func (mock *MockIpam) GetReservations(string) ([]models.Reservation, error) {
	if mock.Err != nil {
		return []models.Reservation{}, mock.Err
	}

	return mock.Reservations, mock.Err
}

// GetReservation ...
func (mock *MockIpam) GetReservation(string) (models.Reservation, error) {
	if mock.Err != nil {
		return models.Reservation{}, mock.Err
	}

	return mock.Reservations[0], mock.Err
}

// CreateReservation ...
func (mock *MockIpam) CreateReservation(reservation models.Reservation) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.ReservationCreated = reservation

	return mock.Err
}

// UpdateReservation ...
func (mock *MockIpam) UpdateReservation(reservation models.Reservation) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.ReservationUpdated = reservation

	return mock.Err
}

// DeleteReservation ...
func (mock *MockIpam) DeleteReservation(id string) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.ReservationDeleted = id

	return mock.Err
}

// GetPoolReservations ...
func (mock *MockIpam) GetPoolReservations(string) ([]models.Reservation, error) {
	if mock.Err != nil {
		return []models.Reservation{}, mock.Err
	}

	return mock.Reservations, mock.Err
}

// GetReservations ...
func (mock *MockIpam) GetLeases(string) ([]models.Lease, error) {
	if mock.Err != nil {
		return []models.Lease{}, mock.Err
	}

	return mock.Leases, mock.Err
}

// GetReservation ...
func (mock *MockIpam) GetLease(string) (models.Lease, error) {
	if mock.Err != nil {
		return models.Lease{}, mock.Err
	}

	return mock.Leases[0], mock.Err
}

// UpdateReservation ...
func (mock *MockIpam) UpdateLease(lease models.Lease) error {
	if mock.Err != nil {
		return mock.Err
	}

	mock.LeaseUpdated = lease

	return mock.Err
}

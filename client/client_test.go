package ipamapi_test

import (
	. "github.com/RackHD/ipam/client"
	"github.com/RackHD/ipam/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client tests", func() {
	var ipamAddress, start, end string
	var err error

	BeforeEach(func() {
		ipamAddress = "127.0.0.1:8000"
		start = "192.168.1.10"
		end = "192.168.1.11"
	})
	var _ = Describe("Pools tests", func() {
		var ipamClient *Client
		var pool resources.PoolV1

		BeforeEach(func() {
			ipamClient = NewClient(ipamAddress)
			Expect(ipamClient).ToNot(BeNil())
		})

		AfterEach(func() {
			poolLocation, err := ipamClient.DeletePool(pool.ID, pool)
			Expect(err).To(BeNil())
			Expect(poolLocation).To(Equal("/pools"))
		})
		It("Should create a pool and return that pool object", func() {

			pool = resources.PoolV1{
				Name:     "Pool1",
				Metadata: "yodawg I heard you like interfaces",
			}

			pool, err = ipamClient.CreateShowPool(pool)
			Expect(err).To(BeNil())
			Expect(pool.ID).ToNot(Equal(""))
			Expect(pool.Name).To(Equal("Pool1"))

		})
	})

	var _ = Describe("Subnets tests", func() {
		var ipamClient *Client
		var pool resources.PoolV1
		var subnet resources.SubnetV1

		BeforeEach(func() {
			ipamClient = NewClient(ipamAddress)
			Expect(ipamClient).ToNot(BeNil())

			pool = resources.PoolV1{
				Name: "SubnetTestPool1",
			}

			pool, err = ipamClient.CreateShowPool(pool)
			Expect(err).To(BeNil())
			Expect(pool.ID).ToNot(Equal(""))
			Expect(pool.Name).To(Equal("SubnetTestPool1"))

		})

		AfterEach(func() {
			poolLocation, err := ipamClient.DeletePool(pool.ID, pool)
			Expect(err).To(BeNil())
			Expect(poolLocation).To(Equal("/pools"))
		})

		It("Should create a subnet and return that subnet object", func() {

			subnet = resources.SubnetV1{
				Name:  "Subnet1",
				Pool:  pool.ID,
				Start: start,
				End:   end,
			}

			subnet, err = ipamClient.CreateShowSubnet(pool.ID, subnet)
			Expect(err).To(BeNil())
			Expect(subnet.ID).ToNot(Equal(""))
			Expect(subnet.Name).To(Equal("Subnet1"))
			Expect(subnet.Pool).To(Equal(pool.ID))

		})

	})
	var _ = Describe("Reservations tests", func() {

		var ipamClient *Client
		var pool resources.PoolV1
		var subnet resources.SubnetV1
		var reservation resources.ReservationV1

		BeforeEach(func() {
			ipamClient = NewClient(ipamAddress)
			Expect(ipamClient).ToNot(BeNil())

			pool = resources.PoolV1{
				Name: "ReservationTestPool1",
			}

			pool, err = ipamClient.CreateShowPool(pool)
			Expect(err).To(BeNil())
			Expect(pool.ID).ToNot(Equal(""))
			Expect(pool.Name).To(Equal("ReservationTestPool1"))

			subnet = resources.SubnetV1{
				Name:  "ReservationTestSubnet1",
				Pool:  pool.ID,
				Start: start,
				End:   end,
			}

			subnet, err = ipamClient.CreateShowSubnet(pool.ID, subnet)
			Expect(err).To(BeNil())
			Expect(subnet.ID).ToNot(Equal(""))
			Expect(subnet.Name).To(Equal("ReservationTestSubnet1"))
			Expect(subnet.Pool).To(Equal(pool.ID))

		})

		AfterEach(func() {
			poolLocation, err := ipamClient.DeletePool(pool.ID, pool)
			Expect(err).To(BeNil())
			Expect(poolLocation).To(Equal("/pools"))
		})

		It("Should create a reservation and return that reservation object", func() {

			reservation = resources.ReservationV1{
				Name:   "Reservation1",
				Subnet: subnet.ID,
			}

			reservation, err = ipamClient.CreateShowReservation(subnet.ID, reservation)
			Expect(err).To(BeNil())
			Expect(reservation.ID).ToNot(Equal(""))
			Expect(reservation.Name).To(Equal("Reservation1"))
			Expect(reservation.Subnet).To(Equal(subnet.ID))

		})

	})
	var _ = Describe("Leases tests", func() {
		var ipamClient *Client
		var pool resources.PoolV1
		var subnet resources.SubnetV1
		var reservation, reservation2 resources.ReservationV1
		var leases, leases2 resources.LeasesV1

		BeforeEach(func() {
			ipamClient = NewClient(ipamAddress)
			Expect(ipamClient).ToNot(BeNil())

			pool = resources.PoolV1{
				Name: "LeaseTestPool1",
			}

			pool, err = ipamClient.CreateShowPool(pool)
			Expect(err).To(BeNil())
			Expect(pool.ID).ToNot(Equal(""))
			Expect(pool.Name).To(Equal("LeaseTestPool1"))

			subnet = resources.SubnetV1{
				Name:  "LeaseTestSubnet1",
				Pool:  pool.ID,
				Start: start,
				End:   end,
			}

			subnet, err = ipamClient.CreateShowSubnet(pool.ID, subnet)
			Expect(err).To(BeNil())
			Expect(subnet.ID).ToNot(Equal(""))
			Expect(subnet.Name).To(Equal("LeaseTestSubnet1"))
			Expect(subnet.Pool).To(Equal(pool.ID))

			reservation = resources.ReservationV1{
				Name:   "LeaseTestReservation1",
				Subnet: subnet.ID,
			}

			reservation, err = ipamClient.CreateShowReservation(subnet.ID, reservation)
			Expect(err).To(BeNil())
			Expect(reservation.ID).ToNot(Equal(""))
			Expect(reservation.Name).To(Equal("LeaseTestReservation1"))
			Expect(reservation.Subnet).To(Equal(subnet.ID))

			reservation2 = resources.ReservationV1{
				Name:   "LeaseTestReservation2",
				Subnet: subnet.ID,
			}

			reservation2, err = ipamClient.CreateShowReservation(subnet.ID, reservation2)
			Expect(err).To(BeNil())
			Expect(reservation2.ID).ToNot(Equal(""))
			Expect(reservation2.Name).To(Equal("LeaseTestReservation2"))
			Expect(reservation2.Subnet).To(Equal(subnet.ID))

		})

		AfterEach(func() {
			poolLocation, err := ipamClient.DeletePool(pool.ID, pool)
			Expect(err).To(BeNil())
			Expect(poolLocation).To(Equal("/pools"))
		})

		It("Should show all leases", func() {
			leases, err = ipamClient.IndexLeases(reservation.ID)
			Expect(err).To(BeNil())
			Expect(leases.Leases[0].ID).ToNot(Equal(""))
			Expect(leases.Leases[0].Reservation).To(Equal(reservation.ID))

			leases2, err = ipamClient.IndexLeases(reservation2.ID)
			Expect(err).To(BeNil())
			Expect(leases2.Leases[0].ID).ToNot(Equal(""))
			Expect(leases2.Leases[0].Reservation).To(Equal(reservation2.ID))

		})
	})
})

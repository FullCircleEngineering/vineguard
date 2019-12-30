package cache

import (
	"sync"
	"time"
	"vineguard/internal/ttndb"

	"github.com/sirupsen/logrus"
)

type Svc struct {
	repo   *repo
	ttnSvc ttndb.Svc
}

func NewService() *Svc {
	repo := NewRepo()
	ttnSvc := ttndb.NewService()
	return &Svc{
		repo:   repo,
		ttnSvc: ttnSvc,
	}
}

func (s *Svc) UpdateData() error {
	// identify all devices linked to our application
	// for each device, retrieve a week of device data, temporally downsample the timeseries to 10 minute averages
	// write the downsampled data to our cache

	deviceList, err := s.ttnSvc.GetDevices()
	if err != nil {
		logrus.Errorf("Problem fetching device list. %s", err)
	}
	var wg sync.WaitGroup
	wg.Add(len(deviceList))
	logrus.Info("Updating cache...")

	for pid, device := range deviceList {
		go func(pid int, device string) {
			defer wg.Done()
			data, err := s.ttnSvc.GetAllFromDevice(device)
			if err != nil {
				logrus.Errorf("Could not fetch data for device %s. Skipping...  %s", device, err)
				return
			}
			_ = s.repo.AddData(device, data)
		}(pid, device)
	}

	wg.Wait()
	return nil
}

func (s *Svc) GetUserRegistrations(usrEmail string) ([]map[string]string, error) {
	return s.repo.GetUserRegistrations(usrEmail)
}

func (s *Svc) CreateNewRegistration(usrEmail string, reg map[string]string) error {
	_ = s.repo.AddRegistration(usrEmail, reg)
	return nil
}

func (s *Svc) GetDeviceData(devId string) (ttndb.LSB50Msgs, error) {
	return ttndb.LSB50Msgs{}, nil
}

func (s *Svc) Run() error {
	for {
		_ = s.UpdateData()
		time.Sleep(time.Second * 60)
	}
	return nil
}

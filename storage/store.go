package storage

import "github.com/masl/answertag/cloud"

type Store interface {
	Create(cloud *cloud.Cloud) error
	Update(cloud *cloud.Cloud) error
	ReadByID(id string) (*cloud.Cloud, error)
}

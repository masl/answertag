package storage

import "github.com/masl/answertag/cloud"

type Store interface {
	Add(cloud *cloud.Cloud) error
	GetById(id string) (*cloud.Cloud, error)
}

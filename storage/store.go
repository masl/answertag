package storage

import "github.com/masl/answertag/cloud"

type Store interface {
	Add(cloud *cloud.Cloud) error
	GetByID(id string) (*cloud.Cloud, error)

	// TODO: add cloud
	AddTag(tag string) error
	GetAllTags() ([]string, error)
}

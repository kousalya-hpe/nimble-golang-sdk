// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Disk shelf and head unit houses disks and controller.
const (
    shelfPath = "shelves"
)

// ShelfObjectSet
type ShelfObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Shelf object
func (objectSet *ShelfObjectSet) CreateObject(payload *model.Shelf) (*model.Shelf, error) {
	response, err := objectSet.Client.Post(shelfPath, payload)
	return response.(*model.Shelf), err
}

// UpdateObject Modify existing Shelf object
func (objectSet *ShelfObjectSet) UpdateObject(id string, payload *model.Shelf) (*model.Shelf, error) {
	response, err := objectSet.Client.Put(shelfPath, id, payload)
	return response.(*model.Shelf), err
}

// DeleteObject deletes the Shelf object with the specified ID
func (objectSet *ShelfObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(shelfPath, id)
}

// GetObject returns a Shelf object with the given ID
func (objectSet *ShelfObjectSet) GetObject(id string) (*model.Shelf, error) {
	response, err := objectSet.Client.Get(shelfPath, id, model.Shelf{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Shelf), err
}

// GetObjectList returns the list of Shelf objects
func (objectSet *ShelfObjectSet) GetObjectList() ([]*model.Shelf, error) {
	response, err := objectSet.Client.List(shelfPath)
	return buildShelfObjectSet(response), err
}

// GetObjectListFromParams returns the list of Shelf objects using the given params query info
func (objectSet *ShelfObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Shelf, error) {
	response, err := objectSet.Client.ListFromParams(shelfPath, params)
	return buildShelfObjectSet(response), err
}

// generated function to build the appropriate response types
func buildShelfObjectSet(response interface{}) ([]*model.Shelf) {
	values := reflect.ValueOf(response)
	results := make([]*model.Shelf, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Shelf{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
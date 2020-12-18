package typesense

import (
	"context"

	"github.com/v-byte-cpu/typesense-go/typesense/api"
)

type DocumentInterface interface {
	Retrieve() (map[string]interface{}, error)
	Update(document interface{}) (map[string]interface{}, error)
	Delete() (map[string]interface{}, error)
}

type document struct {
	apiClient      api.ClientWithResponsesInterface
	collectionName string
	documentID     string
}

func (d *document) Retrieve() (map[string]interface{}, error) {
	response, err := d.apiClient.GetDocumentWithResponse(context.Background(),
		d.collectionName, d.documentID)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &httpError{status: response.StatusCode(), body: response.Body}
	}
	return *response.JSON200, nil
}

func (d *document) Update(document interface{}) (map[string]interface{}, error) {
	response, err := d.apiClient.UpdateDocumentWithResponse(context.Background(),
		d.collectionName, d.documentID, document)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &httpError{status: response.StatusCode(), body: response.Body}
	}
	return *response.JSON200, nil
}

func (d *document) Delete() (map[string]interface{}, error) {
	response, err := d.apiClient.DeleteDocumentWithResponse(context.Background(),
		d.collectionName, d.documentID)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &httpError{status: response.StatusCode(), body: response.Body}
	}
	return *response.JSON200, nil
}

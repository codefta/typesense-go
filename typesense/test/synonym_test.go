//go:build integration
// +build integration

package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typesense/typesense-go/typesense/api/pointer"
)

func TestSearchSynonymRetrieve(t *testing.T) {
	collectionName := createNewCollection(t, "products")
	synonymID := newUUIDName("customize-apple")
	expectedResult := newSearchSynonym(synonymID)

	body := newSearchSynonymSchema()
	_, err := typesenseClient.Collection(collectionName).Synonyms().Upsert(synonymID, body)
	require.NoError(t, err)

	result, err := typesenseClient.Collection(collectionName).Synonym(synonymID).Retrieve()

	require.NoError(t, err)
	expectedResult.Root = pointer.String("")
	require.Equal(t, expectedResult, result)
}

func TestSearchSynonymDelete(t *testing.T) {
	collectionName := createNewCollection(t, "products")
	synonymID := newUUIDName("customize-apple")
	expectedResult := newSearchSynonym(synonymID)

	body := newSearchSynonymSchema()
	_, err := typesenseClient.Collection(collectionName).Synonyms().Upsert(synonymID, body)
	require.NoError(t, err)

	result, err := typesenseClient.Collection(collectionName).Synonym(synonymID).Delete()

	require.NoError(t, err)
	require.Equal(t, expectedResult.Id, result.Id)

	_, err = typesenseClient.Collection(collectionName).Synonym(synonymID).Retrieve()
	require.Error(t, err)
}

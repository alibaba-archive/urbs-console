package dao

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUrbsAcUser(t *testing.T) {
	require := require.New(t)

	_, err := testDaos.UrbsAcUser.Search(context.Background(), "TestUrbsAcUser")
	require.Nil(err)
}

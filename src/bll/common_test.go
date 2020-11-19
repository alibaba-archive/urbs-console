package bll

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/dto"
)

func TestParseGroupUID(t *testing.T) {
	require := require.New(t)

	kind, uid := parseGroupUID("5e43576b62f1d006db395738")
	require.Equal(dto.GroupOrgKind, kind)
	require.Equal("5e43576b62f1d006db395738", uid)

	kind, uid = parseGroupUID("project:5e43576b62f1d006db395738")
	require.Equal("project", kind)
	require.Equal("5e43576b62f1d006db395738", uid)
}

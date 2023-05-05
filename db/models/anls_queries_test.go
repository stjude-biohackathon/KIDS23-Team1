package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_GetAnalysisSuiteByName(t *testing.T) {
	anlsSuite, err := testQueries.GetAnalysisSuiteByName(context.Background(), "test_suite")
	require.NoError(t, err)
	require.Equal(t, 1, anlsSuite.ID)
	require.Equal(t, "test_suite", anlsSuite.SuiteName)
}

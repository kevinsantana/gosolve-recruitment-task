package app

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSliceNums(t *testing.T) {
	testCases := []struct {
		name         string
		fileContent  string
		expectedNums []int64
		expectedErr  error
	}{
		{
			name:         "Valid Input",
			fileContent:  "123 456 789 10\n",
			expectedNums: []int64{123, 456, 789, 10},
		},
		{
			name:        "Invalid Input",
			fileContent: "abc def ghi\n",
			expectedErr: &strconv.NumError{}, // Expect a numerical conversion error
		},
		// {
		// 	name:         "Empty Input",
		// 	fileContent:  "",
		// 	expectedNums: []int64{}, // Expect an empty slice
		// },
		// {
		// 	name:         "Whitespace Input",  // Test case to cover whitespace handling
		// 	fileContent:  "   \n",
		// 	expectedNums: []int64{},
		// },
		{
			name:        "Mixed Valid and Invalid Input", // Handling mixed inputs
			fileContent: "123 abc 456\n",
			expectedErr: &strconv.NumError{}, // Expect a numerical conversion error
		},
		{
			name:         "Negative Numbers",
			fileContent:  "-1 -2 -3 4 5\n",
			expectedNums: []int64{-1, -2, -3, 4, 5},
		}, {
			name:         "Multiple Lines", // Test with multiple lines, including empty ones
			fileContent:  "1 2 3\n\n4 5 6\n",
			expectedNums: []int64{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a temporary file with test data
			tmpFile, err := os.CreateTemp("", "input.txt")
			require.NoError(t, err)
			defer os.Remove(tmpFile.Name())

			_, err = tmpFile.WriteString(tc.fileContent)
			require.NoError(t, err)
			require.NoError(t, tmpFile.Close())

			nums, err := GetSliceNums(context.Background(), tmpFile.Name())

			if tc.expectedErr != nil {
				require.Error(t, err)
				assert.IsType(t, tc.expectedErr, err) //Check Error type
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedNums, nums)
			}
		})
	}

	// Test with a non-existent file
	t.Run("Non-Existent File", func(t *testing.T) {
		_, err := GetSliceNums(context.Background(), "non_existent.txt")
		require.Error(t, err)
	})

}

func TestGetSliceNumsErrorScenarios(t *testing.T) {
	invalidFilePath := "this/file/does/not/exist.txt"

	_, err := GetSliceNums(context.Background(), invalidFilePath)
	require.Error(t, err)

	// Create an unreadable file (optional - depends on OS permissions handling)
	unreadableFile, err := os.CreateTemp("", "unreadable.txt")
	require.NoError(t, err)
	defer os.Remove(unreadableFile.Name())
	require.NoError(t, unreadableFile.Chmod(0000)) // Make it unreadable
	require.NoError(t, unreadableFile.Close())

	_, err = GetSliceNums(context.Background(), unreadableFile.Name())
	require.Error(t, err)

}

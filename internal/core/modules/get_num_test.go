package modules

import (
	"context"
	"testing"
	"time"

	"github.com/kevinsantana/gosolve-recruitment-task/internal/core/share"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockApp ...
type MockApp struct {
	mock.Mock
}

// GetSliceNums ...
func (m *MockApp) GetSliceNums(ctx context.Context, path string) ([]int64, error) {
	args := m.Called(ctx, path)
	return args.Get(0).([]int64), args.Error(1)
}

// SearchModule ...
type SearchModule struct {
	NumberLoader func(ctx context.Context, path string) ([]int64, error)
}

// SearchIndexByValue ...
func (s *SearchModule) SearchIndexByValue(ctx context.Context, value int64) (int, int64, error) {
	nums, err := s.NumberLoader(ctx, share.INPUT_FILE)
	if err != nil {
		return 0, 0, err
	}

	for i := range nums {
		conformation := int64(float64(nums[i]) * share.CONFORMATION)

		if nums[i] == value {
			return i, nums[i], nil
		}

		diff := value - nums[i]
		if diff >= -conformation && diff <= conformation {
			return i, nums[i], nil
		}
	}

	return -1, 0, nil
}

func TestSearchIndexByValue(t *testing.T) {
	testCases := []struct {
		name          string
		value         int64
		mockNums      []int64
		mockErr       error
		expectedIndex int
		expectedValue int64
		expectedError error
	}{
		{
			name:          "Exact Match",
			value:         456,
			mockNums:      []int64{123, 456, 789},
			expectedIndex: 1,
			expectedValue: 456,
		},
		{
			name:          "Match with conformation",
			value:         124,
			mockNums:      []int64{123, 456, 789},
			expectedIndex: 0,
			expectedValue: 123,
		},
		{
			name:          "No Match",
			value:         999,
			mockNums:      []int64{123, 456, 789},
			expectedIndex: -1,
			expectedValue: 0, // Explicitly set expectedValue to 0 when not found.
		},
		{
			name:          "Error Loading Numbers",
			value:         456,
			mockErr:       assert.AnError, // Use assert.AnError for any error.
			expectedError: assert.AnError,
		},
		{
			name:          "Empty Slice",
			value:         456,
			mockNums:      []int64{},
			expectedIndex: -1,
			expectedValue: 0, // Explicitly set expectedValue when the slice is empty.

		},
		{ // Add a test case to handle a zero value input
			name:          "Zero Value Input",
			value:         0,
			mockNums:      []int64{0, 10, 20},
			expectedIndex: 0,
			expectedValue: 0,
		},
		{ // Add a test case where value is negative and there is a negative value that matches
			name:          "Negative Match",
			value:         -10,
			mockNums:      []int64{-10, 10, 20},
			expectedIndex: 0,
			expectedValue: -10,
		},
		{
			name:          "Negative Value, Match With Conformation",
			value:         -11,
			mockNums:      []int64{-10, 10, 20},
			expectedIndex: -1,
			expectedValue: 0,
		}, { // Add test case for empty nums and value is zero
			name:          "Empty Slice Zero Value",
			value:         -1,
			mockNums:      []int64{},
			expectedIndex: -1,
			expectedValue: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockApp := &MockApp{}
			searchModule := &SearchModule{NumberLoader: mockApp.GetSliceNums}

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			mockApp.On("GetSliceNums", ctx, share.INPUT_FILE).Return(tc.mockNums, tc.mockErr).Once()

			actualIndex, actualValue, actualError := searchModule.SearchIndexByValue(ctx, tc.value)

			if tc.expectedError != nil {
				require.Error(t, actualError)
			} else {
				require.NoError(t, actualError)
				assert.Equal(t, tc.expectedIndex, actualIndex)
				assert.Equal(t, tc.expectedValue, actualValue)
			}

			mockApp.AssertExpectations(t)
		})
	}
}

// Example demonstrating usage (testable example)
func ExampleSearchModule_SearchIndexByValue() {
	mockApp := &MockApp{}
	searchModule := &SearchModule{NumberLoader: mockApp.GetSliceNums}

	mockApp.On("GetSliceNums", mock.Anything, share.INPUT_FILE).Return([]int64{10, 20, 30}, nil).Once()

	ctx := context.Background()
	index, value, err := searchModule.SearchIndexByValue(ctx, 20)
	if err != nil {
		// Handle error
	}

	println(index, value)

	mockApp.AssertExpectations(nil) // Assert after printing the output in the example.
}

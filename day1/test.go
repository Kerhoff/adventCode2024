package main

import (
	"reflect"
	"testing"
)

func TestCalculateTotalDistance(t *testing.T) {
	tests := []struct {
		name      string
		leftList  []int
		rightList []int
		expected  int
	}{
		{
			name:      "Example Input",
			leftList:  []int{3, 4, 2, 1, 3, 3},
			rightList: []int{4, 3, 5, 3, 9, 3},
			expected:  11,
		},
		{
			name:      "Same Lists",
			leftList:  []int{1, 2, 3},
			rightList: []int{1, 2, 3},
			expected:  0,
		},
		{
			name:      "Different Lists",
			leftList:  []int{1, 3, 5},
			rightList: []int{2, 4, 6},
			expected:  3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculateTotalDistance(test.leftList, test.rightList)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestReadInputFile(t *testing.T) {
	tests := []struct {
		name          string
		inputContent  string
		expectedLeft  []int
		expectedRight []int
		expectError   bool
	}{
		{
			name:          "Valid Input",
			inputContent:  "1 2\n3 4\n5 6\n",
			expectedLeft:  []int{1, 3, 5},
			expectedRight: []int{2, 4, 6},
			expectError:   false,
		},
		{
			name:          "Invalid Input",
			inputContent:  "1 a\n3 4\n",
			expectedLeft:  nil,
			expectedRight: nil,
			expectError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fileName := "test_input.txt"
			// Write input content to a temporary file
			if err := os.WriteFile(fileName, []byte(test.inputContent), 0644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}
			defer os.Remove(fileName)

			leftList, rightList, err := readInputFile(fileName)
			if test.expectError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(leftList, test.expectedLeft) {
				t.Errorf("Expected left list %v, got %v", test.expectedLeft, leftList)
			}
			if !reflect.DeepEqual(rightList, test.expectedRight) {
				t.Errorf("Expected right list %v, got %v", test.expectedRight, rightList)
			}
		})
	}
}

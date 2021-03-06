// Copyright 2021 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package bytestats

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinCount(t *testing.T) {
	testCases := []struct {
		testBuf  []byte
		expected []int
	}{
		{
			testBuf:  []byte{1, 2, 3, 4, 5},
			expected: []int{0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range testCases {
		count := binCount(tt.testBuf, 256)
		assert.EqualValues(t, tt.expected, count)
	}
}

func TestRollingWindow(t *testing.T) {
	testCases := []struct {
		input    []byte
		window   int
		expected [][]byte
	}{
		{
			input:    []byte{1, 2, 3, 4, 5},
			window:   1,
			expected: [][]byte{{1}, {2}, {3}, {4}, {5}},
		}, {
			input:    []byte{1, 2, 3, 4, 5},
			window:   2,
			expected: [][]byte{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		}, {
			input:    []byte{1, 2, 3, 4, 5},
			window:   3,
			expected: [][]byte{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
		},
	}

	for _, tt := range testCases {
		actual := rollingWindow(tt.input, tt.window)
		assert.EqualValues(t, tt.expected, actual)
	}
}

func TestByteEntropyHistogram(t *testing.T) {
	testCase := []struct {
		testBin  string
		expected []int
	}{
		{
			testBin:  "C:\\Users\\kaplan\\Projects\\saferwall\\binaries\\cmd.exe",
			expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1773, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 1739, 176, 61, 0, 0, 0, 1, 1, 2, 3, 0, 0, 3, 2, 1, 59, 6518, 730, 81, 25, 26, 0, 3, 48, 226, 22, 3, 5, 46, 17, 4, 438, 16926, 430, 290, 194, 325, 69, 156, 228, 816, 528, 110, 107, 155, 210, 108, 1876, 11826, 175, 590, 265, 822, 97, 123, 241, 512, 153, 76, 132, 150, 126, 49, 1047, 7226, 111, 414, 149, 575, 354, 1518, 841, 270, 184, 108, 101, 113, 101, 62, 161, 12585, 284, 1069, 467, 1486, 902, 4674, 2513, 518, 421, 142, 73, 182, 200, 77, 1031, 9367, 763, 1378, 879, 1242, 751, 3924, 2033, 249, 264, 63, 67, 213, 149, 168, 1018, 14417, 1836, 1397, 1153, 1043, 1134, 1966, 1397, 639, 435, 481, 266, 717, 426, 421, 944, 11409, 723, 956, 967, 1437, 1066, 1165, 993, 823, 495, 417, 510, 914, 682, 613, 1406, 3267, 319, 384, 307, 1227, 245, 244, 272, 1182, 118, 132, 219, 1015, 261, 355, 693, 63790, 8277, 14486, 11055, 45079, 8685, 6458, 10308, 42623, 3047, 2740, 5971, 27001, 6994, 11380, 20874, 16247, 2751, 4350, 3416, 11951, 2382, 2421, 3411, 11128, 1062, 909, 2295, 7125, 2242, 3978, 8300, 795, 478, 472, 484, 424, 445, 487, 446, 488, 476, 561, 578, 520, 512, 500, 526},
		},
	}

	step := 1024
	window := 2048

	for _, tt := range testCase {
		bytez, _ := ioutil.ReadFile(tt.testBin)
		vec := byteEntropyHist(bytez, step, window)
		assert.Equal(t, len(tt.expected), len(vec))
		assert.EqualValues(t, tt.expected, vec)
	}
}

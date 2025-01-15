package permos

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	type TestCase struct {
		name     string
		p        fs.FileMode
		expected os.FileMode
	}

	testCases := []TestCase{
		{
			name:     "400",
			p:        FileMode(UserRead),
			expected: fs.FileMode(0400),
		},
		{
			name:     "600",
			p:        FileMode(UserRead, UserWrite),
			expected: fs.FileMode(0600),
		},
		{
			name:     "700",
			p:        FileMode(UserRead, UserWrite, UserExec),
			expected: os.FileMode(0700),
		},
		{
			name:     "040",
			p:        FileMode(GroupRead),
			expected: os.FileMode(0040),
		},
		{
			name:     "060",
			p:        FileMode(GroupRead, GroupWrite),
			expected: os.FileMode(060),
		},
		{
			name:     "070",
			p:        FileMode(GroupRead, GroupWrite, GroupExec),
			expected: os.FileMode(070),
		},
		{
			name: "666",
			p: FileMode(
				UserRead,
				UserWrite,
				GroupRead,
				GroupWrite,
				OtherRead,
				OtherWrite,
			),
			expected: os.FileMode(0666),
		},
		{
			name:     "444",
			p:        FileMode(AllRead),
			expected: os.FileMode(0444),
		},
		{
			name:     "666",
			p:        FileMode(AllRead, AllWrite),
			expected: os.FileMode(0666),
		},
		{
			name:     "777",
			p:        FileMode(AllRead, AllWrite, AllExec),
			expected: os.FileMode(0777),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p)
		})
	}
}

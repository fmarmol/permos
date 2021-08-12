package permos

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	type TestCase struct {
		name     string
		p        Perm
		expected os.FileMode
	}

	testCases := []TestCase{
		{
			name:     "400",
			p:        Perm{UserRead: true},
			expected: os.FileMode(0400),
		},
		{
			name:     "600",
			p:        Perm{UserRead: true, UserWrite: true},
			expected: os.FileMode(0600),
		},
		{
			name:     "700",
			p:        Perm{UserRead: true, UserWrite: true, UserExec: true},
			expected: os.FileMode(0700),
		},
		{
			name:     "040",
			p:        Perm{GroupRead: true},
			expected: os.FileMode(0040),
		},
		{
			name:     "060",
			p:        Perm{GroupRead: true, GroupWrite: true},
			expected: os.FileMode(060),
		},
		{
			name:     "070",
			p:        Perm{GroupRead: true, GroupWrite: true, GroupExec: true},
			expected: os.FileMode(070),
		},
		{
			name: "666",
			p: Perm{
				UserRead:   true,
				UserWrite:  true,
				GroupRead:  true,
				GroupWrite: true,
				OtherRead:  true,
				OtherWrite: true,
			},
			expected: os.FileMode(0666),
		},
		{
			name: "444",
			p: Perm{
				AllRead: true,
			},
			expected: os.FileMode(0444),
		},
		{
			name: "666",
			p: Perm{
				AllRead:  true,
				AllWrite: true,
			},
			expected: os.FileMode(0666),
		},
		{
			name: "777",
			p: Perm{
				AllRead:  true,
				AllWrite: true,
				AllExec:  true,
			},
			expected: os.FileMode(0777),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.FileMode())
		})
	}
}

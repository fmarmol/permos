package permos

import "os"

type Perm struct {
	UserRead   bool
	UserWrite  bool
	UserExec   bool
	GroupRead  bool
	GroupWrite bool
	GroupExec  bool
	AllRead    bool
	AllWrite   bool
	AllExec    bool
}

const (
	AllExec os.FileMode = 1 << iota
	AllWrite
	AllRead
	GroupExec
	GroupWrite
	GroupRead
	UserExec
	UserWrite
	UserRead
)

func (p *Perm) FileMode() (ret os.FileMode) {
	if p.UserRead {
		ret |= UserRead
	}
	if p.UserWrite {
		ret |= UserWrite
	}
	if p.UserExec {
		ret |= UserExec
	}
	if p.GroupRead {
		ret |= GroupRead
	}
	if p.GroupWrite {
		ret |= GroupWrite
	}
	if p.GroupExec {
		ret |= GroupExec
	}
	if p.AllRead {
		ret |= AllRead
	}
	if p.AllWrite {
		ret |= AllWrite
	}
	if p.AllExec {
		ret |= AllExec
	}
	return
}

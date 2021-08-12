package permos

import (
	"io/fs"
)

type Perm struct {
	UserRead   bool
	UserWrite  bool
	UserExec   bool
	GroupRead  bool
	GroupWrite bool
	GroupExec  bool
	OtherRead  bool
	OtherWrite bool
	OtherExec  bool
	AllRead    bool
	AllWrite   bool
	AllExec    bool
}

const (
	OtherExec fs.FileMode = 1 << iota
	OtherWrite
	OtherRead
	GroupExec
	GroupWrite
	GroupRead
	UserExec
	UserWrite
	UserRead
)

const (
	AllRead  fs.FileMode = UserRead | GroupRead | OtherRead
	AllWrite             = UserWrite | GroupWrite | OtherWrite
	AllExec              = UserExec | GroupExec | OtherExec
)

func (p *Perm) FileMode() (ret fs.FileMode) {
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
	if p.OtherRead {
		ret |= OtherRead
	}
	if p.OtherWrite {
		ret |= OtherWrite
	}
	if p.OtherExec {
		ret |= OtherExec
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

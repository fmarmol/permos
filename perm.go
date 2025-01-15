package permos

import (
	"io/fs"
)

type Perm uint32

const (
	OtherExec Perm = 1 << iota
	OtherWrite
	OtherRead
	GroupExec
	GroupWrite
	GroupRead
	UserExec
	UserWrite
	UserRead
	AllRead  = UserRead | GroupRead | OtherRead
	AllWrite = UserWrite | GroupWrite | OtherWrite
	AllExec  = UserExec | GroupExec | OtherExec
)

func FileMode(perms ...Perm) fs.FileMode {
	var ret Perm
	for _, perm := range perms {
		ret |= perm
	}
	return fs.FileMode(ret)

}

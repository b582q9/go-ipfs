//go:build !windows && nofuse
// +build !windows,nofuse

package node

import (
	"errors"

	core "github.com/b582q9/go-ipfs/core"
)

func Mount(node *core.IpfsNode, fsdir, nsdir string) error {
	return errors.New("not compiled in")
}

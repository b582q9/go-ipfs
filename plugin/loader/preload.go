package loader

import (
	pluginbadgerds "github.com/b582q9/go-ipfs/plugin/plugins/badgerds"
	pluginflatfs "github.com/b582q9/go-ipfs/plugin/plugins/flatfs"
	pluginipldgit "github.com/b582q9/go-ipfs/plugin/plugins/git"
	pluginlevelds "github.com/b582q9/go-ipfs/plugin/plugins/levelds"
)

// DO NOT EDIT THIS FILE
// This file is being generated as part of plugin build process
// To change it, modify the plugin/loader/preload.sh

func init() {
	Preload(pluginipldgit.Plugins...)
	Preload(pluginbadgerds.Plugins...)
	Preload(pluginflatfs.Plugins...)
	Preload(pluginlevelds.Plugins...)
}

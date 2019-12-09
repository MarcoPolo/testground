package test

import (
	"context"
	"fmt"

	files "github.com/ipfs/go-ipfs-files"
	utils "github.com/ipfs/testground/plans/chew-large-datasets/utils"
	"github.com/ipfs/testground/sdk/runtime"
)

// IpfsAddDefaults IPFS Add Defaults Test
func IpfsAddDefaults(runenv *runtime.RunEnv) {
	ctx, _ := context.WithCancel(context.Background())
	ipfs, err := utils.CreateIpfsInstance(ctx, nil)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
	}

	err = utils.ForEachCase(runenv, func(unixfsFile files.Node, isDir bool) error {
		t := "file"
		if isDir {
			t = "directory"
		}

		cidFile, err := ipfs.Unixfs().Add(ctx, unixfsFile)
		if err != nil {
			return fmt.Errorf("Could not add %s: %s", t, err)
		}

		fmt.Printf("Added %s to IPFS with CID %s\n", t, cidFile.String())
		return nil
	})

	if err != nil {
		runenv.Abort(err)
		return
	}

	runenv.OK()
}

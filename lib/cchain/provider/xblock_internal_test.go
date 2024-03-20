package provider

import (
	"context"
	"testing"

	"github.com/omni-network/omni/e2e/tutil"
	"github.com/omni-network/omni/lib/cchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestXBlock(t *testing.T) {
	t.Parallel()
	f := fuzz.NewWithSeed(99).NilChance(0).Funcs(
		// Fuzz valid validators.
		func(e *cchain.Validator, c fuzz.Continue) {
			c.FuzzNoCustom(e)
			if e.Power < 0 {
				e.Power = -e.Power
			} else if e.Power == 0 {
				e.Power = 1
			}
		})

	var height uint64
	f.Fuzz(&height)

	valFunc := func(ctx context.Context, h uint64) ([]cchain.Validator, bool, error) {
		require.EqualValues(t, height, h)
		var resp []cchain.Validator
		f.Fuzz(&resp)

		return resp, true, nil
	}
	chainFunc := func(ctx context.Context) (uint64, error) {
		return 77, nil
	}

	prov := Provider{valset: valFunc, chainID: chainFunc}

	block, ok, err := prov.XBlock(context.Background(), height)
	require.NoError(t, err)
	require.True(t, ok)
	tutil.RequireGoldenJSON(t, block)
}
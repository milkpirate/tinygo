//go:build nrf52 || nrf52840 || nrf52833

package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_SPIMode_ApplyTo(t *testing.T) {
	t.Parallel()

	type args struct {
		mode   SPIMode
		config uint32
	}
	type want struct {
		config uint32
	}

	cases := map[string]struct {
		args
		want
	}{
		"Mode0": {
			args: args{
				mode:   SPI_MODE_CPHA0_CPOL0,
				config: 0b_11_0,
			},
			want: want{
				config: 0b_00_0,
			},
		},
		"Mode1": {
			args: args{
				mode:   SPI_MODE_CPHA1_CPOL0,
				config: 0b_10_0,
			},
			want: want{
				config: 0b_01_0,
			},
		},
		"Mode2": {
			args: args{
				mode:   SPI_MODE_CPHA1_CPOL1,
				config: 0b_01_0,
			},
			want: want{
				config: 0b_10_0,
			},
		},
		"Mode3": {
			args: args{
				mode:   SPI_MODE_CPHA0_CPOL1,
				config: 0b_00_0,
			},
			want: want{
				config: 0b_11_0,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actl := tc.args.mode.ApplyTo(tc.args.config)
			require.Equal(t, tc.want.config, actl)
		})
	}
}

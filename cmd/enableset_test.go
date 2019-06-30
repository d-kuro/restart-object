package cmd_test

import (
	"reflect"
	"testing"

	"github.com/d-kuro/restart-object/cmd"
)

var (
	enable = cmd.RestartOptions{
		Enable: []string{"nginx"},
	}
	enableAll = cmd.RestartOptions{
		EnableAll: true,
		Disable:   []string{"nginx"},
	}
	disable = cmd.RestartOptions{
		Enable:  []string{"nginx", "fluentd"},
		Disable: []string{"nginx"},
	}
	disableAll = cmd.RestartOptions{
		Enable:     []string{"nginx", "fluentd", "metrics-server"},
		DisableAll: true,
	}
	disableOnly = cmd.RestartOptions{
		Disable: []string{"nginx"},
	}
)

var cases = []struct {
	name        string
	option      cmd.RestartOptions
	wantEnable  cmd.EnableSet
	wantDisable cmd.DisableSet
}{
	{
		name:        "enable",
		option:      enable,
		wantEnable:  cmd.EnableSet{"nginx"},
		wantDisable: cmd.DisableSet{},
	},
	{
		name:        "enable-all",
		option:      enableAll,
		wantEnable:  cmd.EnableSet{},
		wantDisable: cmd.DisableSet{},
	},
	{
		name:        "disable",
		option:      disable,
		wantEnable:  cmd.EnableSet{"fluentd"},
		wantDisable: cmd.DisableSet{},
	},
	{
		name:        "disable-all",
		option:      disableAll,
		wantEnable:  cmd.EnableSet{"nginx", "fluentd", "metrics-server"},
		wantDisable: cmd.DisableSet{},
	},
	{
		name:        "disable-only",
		option:      disableOnly,
		wantEnable:  cmd.EnableSet{},
		wantDisable: cmd.DisableSet{"nginx"},
	},
}

func TestEnableSetBuild(t *testing.T) {
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			e, d, err := cmd.EnableSetBuild(&tt.option)
			if err != nil {
				t.Errorf("error: %s", err)
			}
			if !reflect.DeepEqual(e, tt.wantEnable) {
				t.Errorf("got enable: %v, want enable: %v", e, tt.wantEnable)
			}
			if !reflect.DeepEqual(d, tt.wantDisable) {
				t.Errorf("got disable: %v, want disable: %v", d, tt.wantDisable)
			}
		})
	}
}

package cmd

import (
	"fmt"
)

func ValidateAllDisableEnableOptions(option *RestartOptions) error {
	if option.EnableAll && option.DisableAll {
		return fmt.Errorf("--enable-all and --disable-all options must not be combined")
	}

	if option.DisableAll {
		if len(option.Enable) == 0 {
			return fmt.Errorf("all linters were disabled, but no one linter was enabled: must enable at least one")
		}

		if len(option.Disable) != 0 {
			return fmt.Errorf("can't combine options --disable-all and --disable %s", option.Disable[0])
		}
	}

	if option.EnableAll && len(option.Enable) != 0 {
		return fmt.Errorf("can't combine options --enable-all and --enable %s", option.Enable[0])
	}

	return nil
}

package cmd

type (
	EnableSet  []string
	DisableSet []string
)

func EnableSetBuild(o *RestartOptions) (EnableSet, DisableSet, error) {
	err := ValidateAllDisableEnableOptions(o)
	if err != nil {
		return nil, nil, err
	}

	switch {
	case o.EnableAll:
		return EnableSet{}, DisableSet{}, nil
	case o.DisableAll:
		return EnableSet(o.Enable), DisableSet{}, nil
	case len(o.Enable) == 0 && len(o.Disable) > 0:
		return EnableSet{}, DisableSet(o.Disable), nil
	case len(o.Enable) > 0 && len(o.Disable) == 0:
		return EnableSet(o.Enable), DisableSet{}, nil
	}

	enableMap := make(map[string]struct{})
	for _, e := range o.Enable {
		if _, ok := enableMap[e]; !ok {
			enableMap[e] = struct{}{}
		}
	}

	for _, d := range o.Disable {
		delete(enableMap, d)
	}

	result := make([]string, 0, len(enableMap))
	for k := range enableMap {
		result = append(result, k)
	}

	return EnableSet(result), DisableSet{}, nil
}

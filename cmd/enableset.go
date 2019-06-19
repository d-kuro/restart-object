package cmd

func EnableSetBuild(option *RestartOptions) ([]string, error) {
	err := ValidateAllDisableEnableOptions(option)
	if err != nil {
		return nil, err
	}

	switch {
	case option.EnableAll:
		return []string{}, nil
	case option.DisableAll:
		return option.Enable, nil
	}

	enableMap := make(map[string]struct{})
	for _, e := range option.Enable {
		if _, ok := enableMap[e]; !ok {
			enableMap[e] = struct{}{}
		}
	}

	for _, d := range option.Disable {
		delete(enableMap, d)
	}

	result := make([]string, 0, len(enableMap))
	for k := range enableMap {
		result = append(result, k)
	}

	return result, nil
}

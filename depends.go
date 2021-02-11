package chopsticks

func installOrder(apps []string, arch string) []string {
	res := []string{}
	for _, app := range apps {

	}

	return res
}

func depResolve(app, arch string, resolved, unresolved []string) {
	app, bucket, _ := parseApp(app)
	unresolved = append(unresolved, app)

}

func runtimeDeps(mf Manifest) []string {
	if *mf.Depends.String != "" {
		return []string{*mf.Depends.String}
	}
	if len(mf.Depends.StringArray) != 0 {
		return mf.Depends.StringArray
	}
	return []string{}
}

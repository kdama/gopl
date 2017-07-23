package golist

// Dependees は、指定されたパッケージに推移的に依存しているパッケージの一覧を返します。
func Dependees(template ...string) ([]Package, error) {
	packages, err := List(template...)
	if err != nil {
		return nil, err
	}
	all, err := List("...")
	if err != nil {
		return nil, err
	}

	var dependees []Package
	for _, dependee := range all {
		depends := false
	loopDependency:
		for _, dependency := range dependee.Deps {
			for _, pack := range packages {
				if dependency == pack.ImportPath {
					depends = true
					break loopDependency
				}
			}
		}
		if depends {
			dependees = append(dependees, dependee)
		}
	}
	return dependees, nil
}

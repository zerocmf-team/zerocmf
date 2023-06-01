package model

type ComponentsTree struct {
	ComponentName string           `json:"componentName"`
	FileName      string           `json:"fileName"`
	Props         interface{}      `json:"props"`
	CSS           string           `json:"css"`
	Children      []ComponentsTree `json:"children"`
}

type Schema struct {
	ComponentsTree []ComponentsTree `json:"componentsTree"`
}

func FindComponents(components []ComponentsTree, targetComponentName string) (result []ComponentsTree) {
	for _, component := range components {
		if component.ComponentName == targetComponentName {
			result = append(result, component)
		}

		if len(component.Children) > 0 {
			children := FindComponents(component.Children, targetComponentName)
			result = append(result, children...)
		}
	}
	return
}

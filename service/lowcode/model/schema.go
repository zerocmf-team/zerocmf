package model

type ComponentProps struct {
	Main   bool     `json:"main,omitempty"`
	Label  string   `json:"label"`
	Name   string   `json:"name"`
	Unique bool     `json:"unique"`
	Rules  []SRules `json:"rules,omitempty"`
}

type ComponentsTree struct {
	ComponentName string           `json:"componentName"`
	FileName      string           `json:"fileName"`
	Props         ComponentProps   `json:"props"`
	CSS           string           `json:"css"`
	Children      []ComponentsTree `json:"children"`
}

type Schema struct {
	ComponentsTree []ComponentsTree `json:"componentsTree"`
}

func FindComponents(components []ComponentsTree, targetComponentName string) (result []ComponentsTree) {
	for _, component := range components {

		if component.ComponentName == "Form" && !component.Props.Main {
			continue
		}

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

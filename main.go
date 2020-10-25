package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type UseCase struct {
	Version                 string
	Name                    string   `yaml:"name"`
	RelatedRequirements     []string `yaml:"related_requirements"`
	GoalInContext           []string `yaml:"goal_in_context"`
	Preconditions           []string `yaml:"preconditions"`
	SuccessfulEndConditions []string `yaml:"successful_end_conditions"`
	FailedEndConditions     []string `yaml:"failed_end_conditions"`
	PrimaryActors           []string `yaml:"primary_actors"`
	SecondaryActors         []string `yaml:"secondary_actors"`
	Trigger                 []string `yaml:"trigger"`
	MainFlow                []string `yaml:"main_flow"`
	Extensions              []string `yaml:"extensions"`
}

const readme = `
{{.Name}}

| Title | Description |
| -     | -           |
| Name | {{ .Name }} |
| Related requirements | {{ list .RelatedRequirements }}|
| Goal in Context | {{ list .GoalInContext }}|
| Preconditions | {{ list .Preconditions }}|
| Successful end conditions | {{ list .SuccessfulEndConditions }}|
| Failed end conditions | {{ list .FailedEndConditions }}|
| Primary Actors | {{ list .PrimaryActors }}|
| Secondary Actors | {{ list .SecondaryActors }}|
| Trigger | {{ list .Trigger }}|
| Main flow | {{ list .MainFlow }}|
| Extensions | {{ list .Extensions }}|
`

func main() {
	b, err := ioutil.ReadFile("usecase.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var u UseCase
	if err := yaml.Unmarshal(b, &u); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", u)
	t := template.Must(template.New("README.md").Funcs(template.FuncMap{
		"list": func(str []string) template.HTML {
			if len(str) == 1 {
				return template.HTML(str[0])
			}
			//result := make([]string, len(str))
			//for key, value := range str {
			//result[key] = fmt.Sprintf("%d. %s", key+1, value)
			//}
			out := strings.Join(str, "<br/>")
			return template.HTML(out)
		},
	}).Parse(readme))

	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := t.Execute(f, u); err != nil {
		log.Fatal(err)
	}
}

package main

import "fmt"
import "strings"
import "github.com/xlab/treeprint"
import "github.com/Knetic/govaluate"
import "github.com/muhqu/go-gherkin"
import "github.com/muhqu/go-gherkin/nodes"

func EvaluateExpression(expression string, parameters map[string]interface{}) (interface{}, error) {
	var err error
	var res interface{}
	var exp *govaluate.EvaluableExpression
	if exp, err = govaluate.NewEvaluableExpression(expression); err == nil {
		res, err = exp.Evaluate(parameters)
		return res, err
	}
	return nil, err
}

func ParseGiven(step nodes.StepNode, params map[string]interface{}, example interface{}) map[string]interface{} {
	var indexMap map[string]int = make(map[string]int)
	params["example"] = example

	if table := step.Table(); table != nil {
		rows := table.Rows()
		rowLen := len(rows)
		if rowLen > 1 {
			commandType := strings.TrimSpace(strings.ToLower(step.Text()))
			for i, rowIndexes := range rows[0] {
				indexMap[rowIndexes] = i + 1
			}
			switch commandType {
			case GIVEN_CMD_DEFINE:
				for r := 1; r < rowLen; r++ {
					row := rows[r]
					name, instance := GetInstanceFor(indexMap, row)
					params[name] = instance
				}
			}
		}
	}

	return params
}

func ParseStep(step nodes.StepNode, params map[string]interface{}) (interface{}, error) {
	stepText := step.Text()
	method := index[stepText]

	if method != nil {
		return method(), nil
	} else {
		result, err := EvaluateExpression(stepText, params)
		return result, err
	}
}

func LoadFeature(featureDef string) {
	var err error
	var feature nodes.FeatureNode
	if feature, err = gherkin.ParseGherkinFeature(featureDef); err != nil {
		return
	}

	tree := treeprint.New()
	params := make(map[string]interface{})

scenarioLoop:
	for _, scenario := range feature.Scenarios() {

		scenarioTags := scenario.Tags()
		for _, tag := range scenarioTags {
			tagProps := strings.Split(tag, "::")
			for _, prop := range tagProps {
				if prop == "Rollback" {
					continue scenarioLoop
				}
			}
		}

		var name string
		var example interface{}
		var examples nodes.OutlineExamplesNode
		var indexMap map[string]int = make(map[string]int)

		exampleIndex, exampleCount := 1, 0

		if outline, ok := scenario.(nodes.OutlineNode); ok {
			examples = outline.Examples()
			exampleCount = len(examples.Table().Rows()) - 1
		}

		if exampleCount > 0 {
			rows := examples.Table().Rows()
			for i, rowIndexes := range rows[0] {
				indexMap[rowIndexes] = i + 1
			}
		}

	runOutline:
		var featureBranch treeprint.Tree
		ftitle, stitle := feature.Title(), scenario.Title()
		if exampleCount > 0 {
			exampleData := examples.Table().Rows()[exampleIndex]
			name, example = GetInstanceFor(indexMap, exampleData)
			featureBranch = tree.AddBranch(fmt.Sprintf("Feature <%s>: %s", name, ftitle))
		} else {
			featureBranch = tree.AddBranch(fmt.Sprintf("Feature: %s", ftitle))
		}
		scenarioBranch := featureBranch.AddBranch(fmt.Sprintf("Scenario: %s", stitle))

	stepIteration:
		for _, step := range scenario.Steps() {
			switch step.StepType() {
			case "Given":
				params = ParseGiven(step, params, example)
			case "Then", "When", "And", "Or", "But":
				r, err := ParseStep(step, params)
				scenarioBranch.AddNode(fmt.Sprintf("Step: %s", step.Text()))

				if err != nil {
					scenarioBranch.AddNode(fmt.Sprintf("Fail: %v", err))
					break stepIteration
				}

				switch result := r.(type) {
				case bool:
					if !result {
						scenarioBranch.AddNode("Fail: Step returned false")
						break stepIteration
					}
				}
			}
		}

		if exampleIndex < exampleCount {
			exampleIndex += 1
			goto runOutline
		}

	}

	fmt.Println(tree.String())
}

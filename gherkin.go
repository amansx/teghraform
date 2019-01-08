package main

import "fmt"
import "strings"
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

func ParseGiven(step nodes.StepNode) map[string]interface{} {
	var objects map[string]interface{} = make(map[string]interface{})
	var indexMap map[string]int = make(map[string]int)

	if table := step.Table(); table != nil {
		rows := table.Rows()
		rowLen := len(rows)
		if rowLen > 1 {
			instanceType := strings.TrimSpace(strings.ToLower(step.Text()))
			for i, rowIndexes := range rows[0] {
				indexMap[rowIndexes] = i + 1
			}
			for r := 1; r < rowLen; r++ {
				row := rows[r]
				name, instance := GetInstanceFor(instanceType, indexMap, row)
				objects[name] = instance
			}
		}
	}

	return objects
}

func ParseStep(step nodes.StepNode, params map[string]interface{}) (interface{}, error) {
	stepText := step.Text()
	method := index[stepText]

	if method != nil {
		return method(), nil
	} else {
		result, err := EvaluateExpression(stepText, params)
		fmt.Println(stepText, "-->", result)
		return result, err
	}
}

func LoadFeature(featureDef string) {
	var err error
	var feature nodes.FeatureNode
	if feature, err = gherkin.ParseGherkinFeature(featureDef); err != nil {
		fmt.Println(err)
		return
	}

	params := make(map[string]interface{})
	for _, scenario := range feature.Scenarios() {
	stepIteration:
		for _, step := range scenario.Steps() {
			switch step.StepType() {
			case "Given":
				params = ParseGiven(step)
			case "Then", "When", "And", "Or", "But":
				_, err := ParseStep(step, params)
				if err != nil {
					fmt.Println("Step broke rolling back..", err)
					break stepIteration
				}
			}
		}
	}
}

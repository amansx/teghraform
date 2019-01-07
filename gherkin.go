package main

import "fmt"
import "github.com/Knetic/govaluate"
import "github.com/muhqu/go-gherkin"
import "github.com/muhqu/go-gherkin/nodes"

func EvaluateExpression(expression string, parameters map[string]interface{}) (interface{}, error){
	var err error
	var res interface{}
	var exp *govaluate.EvaluableExpression
	if exp, err = govaluate.NewEvaluableExpression(expression); err == nil {
		res, err = exp.Evaluate(parameters);
		return res, err
	}
	return nil, err
}

func ParseGiven(){
	rows := step.Table().Rows()
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
		
		for _, step := range scenario.Steps() {
			
			switch step.StepType() {
				case "Given":
				default:
					v, e := EvaluateExpression(step.Text(), params)
			}
		}
	}
}
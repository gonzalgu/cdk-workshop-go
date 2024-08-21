package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PipelineStackProps struct {
	awscdk.StackProps
}

func NewPipelineStack(scope constructs.Construct, id string, props *PipelineStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	repo := pipelines.CodePipelineSource_GitHub(
		jsii.String("gonzalgu/cdk-workshop-go"),
		jsii.String("main"),
		nil,
	)
	pipelines.NewCodePipeline(stack, jsii.String("Pipeline"), &pipelines.CodePipelineProps{
		PipelineName: jsii.String("WorkshopPipeline"),
		Synth: pipelines.NewShellStep(jsii.String("Synth"), &pipelines.ShellStepProps{
			Input: repo,
			Commands: jsii.Strings(
				"npm install -g aws-cdk",
				"goenv install 1.22.3",
				"goenv local 1.22.3",
				"npx cdk synth",
			),
		}),
	})
	return stack
}

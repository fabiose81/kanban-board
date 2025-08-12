package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"kanban-board/structs"
	"kanban-board/utils"
)

func Save(board structs.BoardRequest) (*structs.BoardResponseSave, error) {
	payload, err := json.Marshal(board)
	if err != nil {
		return nil, fmt.Errorf(utils.MSG_MARSHAL_PAYLOAD, err)
	}

	result, err := InvokeLambdaFunction(payload, "AWS_LAMBDA_FUNCTION_SAVE")

	if err != nil {
		return nil, fmt.Errorf(utils.MSG_FAILED_INVOKE_LAMBDA, err)
	}

	var response structs.BoardResponseSave
	if err := json.Unmarshal(result.Payload, &response); err != nil {
		return nil, fmt.Errorf(utils.MSG_FAILED_UNMARSHAL_PAYLOAD, err)
	}

	return &response, nil
}

func Get(userid string) (*structs.BoardResponseGet, error) {
	payload, err := json.Marshal(map[string]string{
		"userid": userid,
	})

	if err != nil {
		return nil, fmt.Errorf(utils.MSG_MARSHAL_PAYLOAD, err)
	}

	result, err := InvokeLambdaFunction(payload, "AWS_LAMBDA_FUNCTION_GET")

	if err != nil {
		return nil, fmt.Errorf(utils.MSG_FAILED_INVOKE_LAMBDA, err)
	}

	var boards structs.BoardResponseGet
	if err := json.Unmarshal(result.Payload, &boards); err != nil {
		return nil, fmt.Errorf(utils.MSG_FAILED_UNMARSHAL_PAYLOAD, err)
	}

	return &boards, nil
}

func InvokeLambdaFunction(payload []byte, functionName string) (*lambda.InvokeOutput, error) {
	ctx := context.TODO()

	cfg, err := LoadCustomAWSConfig(ctx)

	if err != nil {
		return nil, fmt.Errorf(utils.MSG_FAILED_LOAD_AWS, err)
	}

	client := lambda.NewFromConfig(cfg)

	return client.Invoke(ctx, &lambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv(functionName)),
		Payload:        payload,
		InvocationType: types.InvocationTypeRequestResponse,
	})
}

func LoadCustomAWSConfig(ctx context.Context) (aws.Config, error) {
	profile := os.Getenv("AWS_PROFILE")
	region := os.Getenv("AWS_REGION")

	return config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
}

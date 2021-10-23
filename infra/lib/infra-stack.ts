import * as cdk from '@aws-cdk/core';
import * as apigateway from '@aws-cdk/aws-apigateway'
import * as lambda from '@aws-cdk/aws-lambda'
import * as path from "path"


export class InfraStack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);

        const helloWorldLambda = new lambda.Function(this, "helloWorldLambda", {
			runtime: lambda.Runtime.GO_1_X,
			handler: 'helloworld',
			code: lambda.Code.fromAsset(
				path.join(__dirname, "../../lambda/bin")
			),
		});

        const api = new apigateway.RestApi(this, "go-web-api", {
            restApiName: "Go Web API",
        });

        const helloWorldIntegration = new apigateway.LambdaIntegration(helloWorldLambda, {
            requestTemplates: { "application/json": '{ "statusCode": "200" }' }
        });

        api.root.addMethod("GET", helloWorldIntegration); // GET /
    }
}

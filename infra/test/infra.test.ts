import { expect as expectCDK, haveResource } from '@aws-cdk/assert';
import * as cdk from '@aws-cdk/core';
import * as Infra from '../lib/infra-stack';

test('stack includes gateway', () => {
    const app = new cdk.App();
    const stack = new Infra.InfraStack(app, 'MyTestStack');
    expectCDK(stack).to(haveResource('AWS::ApiGateway::RestApi', {
        restApiName: "Go Web Api"
    }))
});

test('stack includes lambda function', () => {
    const app = new cdk.App();
    const stack = new Infra.InfraStack(app, 'MyTestStack');
    expectCDK(stack).to(haveResource('AWS::Lambda::Function'))
});

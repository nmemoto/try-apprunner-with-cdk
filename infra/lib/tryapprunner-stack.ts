import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as apprunner from '@aws-cdk/aws-apprunner-alpha'
import * as ecr from 'aws-cdk-lib/aws-ecr';
export class TryapprunnerStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);
    new apprunner.Service(this, 'Service', {
      source: apprunner.Source.fromEcr({
        imageConfiguration: { 
          port: 8080,
          environment: {
            "PORT": "8080"
          }
        },
        repository: ecr.Repository.fromRepositoryName(this, 'tryapprunnerRepository', 'tryapprunner-ecr'),
        tagOrDigest: 'latest',
      }),
    });

  }
}

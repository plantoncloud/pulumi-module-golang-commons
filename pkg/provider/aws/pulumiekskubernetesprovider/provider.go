package pulumiekskubernetesprovider

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/awscredential"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithCreatedEksClusterWithAwsCredentials returns kubernetes provider for the added eks cluster based on the aws provider
func GetWithCreatedEksClusterWithAwsCredentials(ctx *pulumi.Context, createdEksCluster *eks.Cluster,
	awsCredential *awscredential.AwsCredential,
	dependencies []pulumi.Resource, providerName string) (*kubernetes.Provider, error) {
	provider, err := kubernetes.NewProvider(ctx,
		providerName,
		&kubernetes.ProviderArgs{
			EnableServerSideApply: pulumi.Bool(true),
			Kubeconfig: pulumi.Sprintf(AwsExecPluginKubeConfigTemplate,
				createdEksCluster.Endpoint,
				createdEksCluster.CertificateAuthority.Data().Elem(),
				awsCredential.Spec.AccessKeyId,
				awsCredential.Spec.SecretAccessKey,
				awsCredential.Spec.Region,
			),
		}, pulumi.DependsOn(dependencies))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}

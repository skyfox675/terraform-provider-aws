// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package grafana

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	managedgrafana_sdkv1 "github.com/aws/aws-sdk-go/service/managedgrafana"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceWorkspace,
			TypeName: "aws_grafana_workspace",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLicenseAssociation,
			TypeName: "aws_grafana_license_association",
		},
		{
			Factory:  ResourceRoleAssociation,
			TypeName: "aws_grafana_role_association",
		},
		{
			Factory:  ResourceWorkspace,
			TypeName: "aws_grafana_workspace",
			Name:     "Workspace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceWorkspaceAPIKey,
			TypeName: "aws_grafana_workspace_api_key",
		},
		{
			Factory:  ResourceWorkspaceSAMLConfiguration,
			TypeName: "aws_grafana_workspace_saml_configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Grafana
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*managedgrafana_sdkv1.ManagedGrafana, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return managedgrafana_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}

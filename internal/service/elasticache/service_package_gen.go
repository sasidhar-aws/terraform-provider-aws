// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceReservedCacheNodeOffering,
			Name:    "Reserved Cache Node Offering",
		},
		{
			Factory: newDataSourceServerlessCache,
			Name:    "Serverless Cache",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceReservedCacheNode,
			Name:    "Reserved Cache Node",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory: newServerlessCacheResource,
			Name:    "Serverless Cache",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_elasticache_cluster",
			Name:     "Cluster",
		},
		{
			Factory:  dataSourceReplicationGroup,
			TypeName: "aws_elasticache_replication_group",
			Name:     "Replication Group",
		},
		{
			Factory:  dataSourceSubnetGroup,
			TypeName: "aws_elasticache_subnet_group",
			Name:     "Subnet Group",
		},
		{
			Factory:  dataSourceUser,
			TypeName: "aws_elasticache_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCluster,
			TypeName: "aws_elasticache_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGlobalReplicationGroup,
			TypeName: "aws_elasticache_global_replication_group",
			Name:     "Global Replication Group",
		},
		{
			Factory:  resourceParameterGroup,
			TypeName: "aws_elasticache_parameter_group",
			Name:     "Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceReplicationGroup,
			TypeName: "aws_elasticache_replication_group",
			Name:     "Replication Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSubnetGroup,
			TypeName: "aws_elasticache_subnet_group",
			Name:     "Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_elasticache_user",
			Name:     "User",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUserGroup,
			TypeName: "aws_elasticache_user_group",
			Name:     "User Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUserGroupAssociation,
			TypeName: "aws_elasticache_user_group_association",
			Name:     "User Group Association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ElastiCache
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticache.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return elasticache.NewFromConfig(cfg,
		elasticache.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

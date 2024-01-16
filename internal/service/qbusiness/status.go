// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/qbusiness"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func statusAppAvailability(ctx context.Context, conn *qbusiness.QBusiness, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindAppByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}

func statusIndexAvailability(ctx context.Context, conn *qbusiness.QBusiness, index_id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindIndexByID(ctx, conn, index_id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}

func statusRetrieverAvailability(ctx context.Context, conn *qbusiness.QBusiness, retriever_id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindRetrieverByID(ctx, conn, retriever_id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the details of a secret. It does not include the encrypted secret
// value. Secrets Manager only returns fields that have a value in the response.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:DescribeSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func (c *Client) DescribeSecret(ctx context.Context, params *DescribeSecretInput, optFns ...func(*Options)) (*DescribeSecretOutput, error) {
	if params == nil {
		params = &DescribeSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSecret", params, optFns, c.addOperationDescribeSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecretInput struct {

	// The ARN or name of the secret.
	//
	// For an ARN, we recommend that you specify a complete ARN rather than a partial
	// ARN. See [Finding a secret from a partial ARN].
	//
	// [Finding a secret from a partial ARN]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen
	//
	// This member is required.
	SecretId *string

	noSmithyDocumentSerde
}

type DescribeSecretOutput struct {

	// The ARN of the secret.
	ARN *string

	// The date the secret was created.
	CreatedDate *time.Time

	// The date the secret is scheduled for deletion. If it is not scheduled for
	// deletion, this field is omitted. When you delete a secret, Secrets Manager
	// requires a recovery window of at least 7 days before deleting the secret. Some
	// time after the deleted date, Secrets Manager deletes the secret, including all
	// of its versions.
	//
	// If a secret is scheduled for deletion, then its details, including the
	// encrypted secret value, is not accessible. To cancel a scheduled deletion and
	// restore access to the secret, use RestoreSecret.
	DeletedDate *time.Time

	// The description of the secret.
	Description *string

	// The key ID or alias ARN of the KMS key that Secrets Manager uses to encrypt the
	// secret value. If the secret is encrypted with the Amazon Web Services managed
	// key aws/secretsmanager , this field is omitted. Secrets created using the
	// console use an KMS key ID.
	KmsKeyId *string

	// The date that the secret was last accessed in the Region. This field is omitted
	// if the secret has never been retrieved in the Region.
	LastAccessedDate *time.Time

	// The last date and time that this secret was modified in any way.
	LastChangedDate *time.Time

	// The last date and time that Secrets Manager rotated the secret. If the secret
	// isn't configured for rotation or rotation has been disabled, Secrets Manager
	// returns null.
	LastRotatedDate *time.Time

	// The name of the secret.
	Name *string

	// The next rotation is scheduled to occur on or before this date. If the secret
	// isn't configured for rotation or rotation has been disabled, Secrets Manager
	// returns null. If rotation fails, Secrets Manager retries the entire rotation
	// process multiple times. If rotation is unsuccessful, this date may be in the
	// past.
	//
	// This date represents the latest date that rotation will occur, but it is not an
	// approximate rotation date. In some cases, for example if you turn off automatic
	// rotation and then turn it back on, the next rotation may occur much sooner than
	// this date.
	NextRotationDate *time.Time

	// The ID of the service that created this secret. For more information, see [Secrets managed by other Amazon Web Services services].
	//
	// [Secrets managed by other Amazon Web Services services]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/service-linked-secrets.html
	OwningService *string

	// The Region the secret is in. If a secret is replicated to other Regions, the
	// replicas are listed in ReplicationStatus .
	PrimaryRegion *string

	// A list of the replicas of this secret and their status:
	//
	//   - Failed , which indicates that the replica was not created.
	//
	//   - InProgress , which indicates that Secrets Manager is in the process of
	//   creating the replica.
	//
	//   - InSync , which indicates that the replica was created.
	ReplicationStatus []types.ReplicationStatusType

	// Specifies whether automatic rotation is turned on for this secret. If the
	// secret has never been configured for rotation, Secrets Manager returns null.
	//
	// To turn on rotation, use RotateSecret. To turn off rotation, use CancelRotateSecret.
	RotationEnabled *bool

	// The ARN of the Lambda function that Secrets Manager invokes to rotate the
	// secret.
	RotationLambdaARN *string

	// The rotation schedule and Lambda function for this secret. If the secret
	// previously had rotation turned on, but it is now turned off, this field shows
	// the previous rotation schedule and rotation function. If the secret never had
	// rotation turned on, this field is omitted.
	RotationRules *types.RotationRulesType

	// The list of tags attached to the secret. To add tags to a secret, use TagResource. To
	// remove tags, use UntagResource.
	Tags []types.Tag

	// A list of the versions of the secret that have staging labels attached.
	// Versions that don't have staging labels are considered deprecated and Secrets
	// Manager can delete them.
	//
	// Secrets Manager uses staging labels to indicate the status of a secret version
	// during rotation. The three staging labels for rotation are:
	//
	//   - AWSCURRENT , which indicates the current version of the secret.
	//
	//   - AWSPENDING , which indicates the version of the secret that contains new
	//   secret information that will become the next current version when rotation
	//   finishes.
	//
	// During rotation, Secrets Manager creates an AWSPENDING version ID before
	//   creating the new secret version. To check if a secret version exists, call GetSecretValue.
	//
	//   - AWSPREVIOUS , which indicates the previous current version of the secret.
	//   You can use this as the last known good version.
	//
	// For more information about rotation and staging labels, see [How rotation works].
	//
	// [How rotation works]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html
	VersionIdsToStages map[string][]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSecret{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSecret"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecret(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSecret",
	}
}

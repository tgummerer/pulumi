// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package simpleinvoke

import (
	"context"
	"reflect"

	"example.com/pulumi-simple-invoke/sdk/go/v10/simpleinvoke/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func SecretInvoke(ctx *pulumi.Context, args *SecretInvokeArgs, opts ...pulumi.InvokeOption) (*SecretInvokeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv SecretInvokeResult
	err := ctx.Invoke("simple-invoke:index:secretInvoke", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type SecretInvokeArgs struct {
	SecretResponse bool   `pulumi:"secretResponse"`
	Value          string `pulumi:"value"`
}

type SecretInvokeResult struct {
	Response string `pulumi:"response"`
	Secret   bool   `pulumi:"secret"`
}

func SecretInvokeOutput(ctx *pulumi.Context, args SecretInvokeOutputArgs, opts ...pulumi.InvokeOption) SecretInvokeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SecretInvokeResultOutput, error) {
			args := v.(SecretInvokeArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv SecretInvokeResult
			secret, err := ctx.InvokePackageRaw("simple-invoke:index:secretInvoke", args, &rv, "", opts...)
			if err != nil {
				return SecretInvokeResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(SecretInvokeResultOutput)
			if secret {
				return pulumi.ToSecret(output).(SecretInvokeResultOutput), nil
			}
			return output, nil
		}).(SecretInvokeResultOutput)
}

type SecretInvokeOutputArgs struct {
	SecretResponse pulumi.BoolInput   `pulumi:"secretResponse"`
	Value          pulumi.StringInput `pulumi:"value"`
}

func (SecretInvokeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretInvokeArgs)(nil)).Elem()
}

type SecretInvokeResultOutput struct{ *pulumi.OutputState }

func (SecretInvokeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretInvokeResult)(nil)).Elem()
}

func (o SecretInvokeResultOutput) ToSecretInvokeResultOutput() SecretInvokeResultOutput {
	return o
}

func (o SecretInvokeResultOutput) ToSecretInvokeResultOutputWithContext(ctx context.Context) SecretInvokeResultOutput {
	return o
}

func (o SecretInvokeResultOutput) Response() pulumi.StringOutput {
	return o.ApplyT(func(v SecretInvokeResult) string { return v.Response }).(pulumi.StringOutput)
}

func (o SecretInvokeResultOutput) Secret() pulumi.BoolOutput {
	return o.ApplyT(func(v SecretInvokeResult) bool { return v.Secret }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretInvokeResultOutput{})
}
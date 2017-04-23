// Copyright © 2017 Toru Motchie MOCHIDA
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package s3

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetWebSite returns the website configuration for a bucket.
func GetWebSite(bucket string) error {
	params := &s3.GetBucketWebsiteInput{
		Bucket: aws.String(bucket),
	}

	resp, err := s3Clinet.GetBucketWebsite(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			return errors.New(awsErr.Message())
		}
		return errors.New(err.Error())
	}

	if resp.IndexDocument != nil {
		fmt.Printf("IndexDocument : %s\n", *resp.IndexDocument.Suffix)
		if resp.ErrorDocument != nil {
			fmt.Printf("ErrorDocument : %s\n", *resp.ErrorDocument.Key)
		}
	} else {
		if resp.RedirectAllRequestsTo != nil {
			fmt.Printf("Host : %s\n", *resp.RedirectAllRequestsTo.HostName)
			fmt.Printf("Protocol : %s\n", *resp.RedirectAllRequestsTo.Protocol)
		}
	}

	return nil
}

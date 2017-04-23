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

package cmd

import (
	"fmt"
	"os"

	"github.com/motchie/s3w/s3"
	"github.com/spf13/cobra"
)

var setBucketPolicyCmd = &cobra.Command{
	Use:   "set-bucket-policy-for-website",
	Short: "Replaces a policy on a bucket for Web Site.",
	Long:  `Replaces a policy on a bucket for Web Site.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(bucket) == 0 {
			fmt.Fprintf(os.Stderr, "Bucketname Empty\n")
			os.Exit(-1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := s3.SetBucketPolicyForWebSite(bucket)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(setBucketPolicyCmd)
	setBucketPolicyCmd.PersistentFlags().StringVar(&bucket, "bucket", "", "The name of S3 bucket to get.")
}

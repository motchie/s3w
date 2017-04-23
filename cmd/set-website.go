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

var (
	flagindex string
	flagerror string
)

var setWebsiteCmd = &cobra.Command{
	Use:   "set-website",
	Short: "Set the website configuration for a bucket.",
	Long: `Set the website configuration for a bucket.
	
	To redirect all requests to another URL or bucket, use set-redirect command.
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(bucket) == 0 {
			fmt.Fprintf(os.Stderr, "Bucketname Empty\n")
			os.Exit(-1)
		}

		if len(flagindex) == 0 {
			fmt.Fprintf(os.Stderr, "IndexDocument Empty\n")
			os.Exit(-1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		param := &s3.SetWebSiteInput{
			IndexDocument: flagindex,
			ErrorDocument: flagerror,
		}

		if err := s3.SetWebSite(bucket, param); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(setWebsiteCmd)
	setWebsiteCmd.Flags().StringVar(&bucket, "bucket", "", "The name of S3 bucket to set the website configuration.")
	setWebsiteCmd.Flags().StringVar(&flagindex, "indexdocument", "", "(Required) A suffix that is appended to a request that is for a directory on the website endpoint.")
	setWebsiteCmd.Flags().StringVar(&flagerror, "errordocument", "", "The object key name to use when a 4XX class error occurs. This key identifies the page that is returned when such an error occurs.")
}

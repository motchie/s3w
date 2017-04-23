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
	"net/url"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/motchie/s3w/s3"
	"github.com/spf13/cobra"
)

var (
	flagurl   string
	targeturl *url.URL
)

var setRedirectAllRequestsCmd = &cobra.Command{
	Use:   "set-redirect-all-requests",
	Short: "Set to redirect all website requests sent to the bucket's endpoint",
	Long:  `Set to redirect all website requests sent to the bucket's endpoint.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(bucket) == 0 {
			fmt.Fprintf(os.Stderr, "Bucketname Empty\n")
			os.Exit(-1)
		}

		if !govalidator.IsURL(flagurl) {
			fmt.Fprintf(os.Stderr, "Invalid URL\n")
			os.Exit(-1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		targeturl, err := url.Parse(flagurl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid URL\n")
			os.Exit(-1)
		}
		if err := s3.SetRedirectAllRequestsTo(bucket, targeturl); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(setRedirectAllRequestsCmd)
	setRedirectAllRequestsCmd.Flags().StringVar(&bucket, "bucket", "", "The name of S3 bucket to set the website configuration.")
	setRedirectAllRequestsCmd.Flags().StringVar(&flagurl, "url", "", "The URL where requests will be redirected. Only scheme and host part of URL will be accepted.")
}

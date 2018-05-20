// Copyright © 2018 Ken'ichiro Oyama <k1lowxb@gmail.com>
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
	"github.com/k1LoW/tbls/db"
	"github.com/k1LoW/tbls/output/md"
	"github.com/spf13/cobra"
	"os"
)

// outputPath is path to generate document
var outputPath string

// force is a flag on whether to force genarate
var force bool

// docCmd represents the doc command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "document a database.",
	Long:  `'tbls doc' analyze a database and generate document in GitHub Friendly Markdown format.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := db.Analyze(dsn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = md.Output(s, outputPath, force)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
	docCmd.Flags().StringVarP(&dsn, "dsn", "u", "", "URL like DSN. ex. postgres://user:pass@localhost/dbname")
	docCmd.Flags().StringVarP(&outputPath, "output", "o", ".", "output filepath")
	docCmd.Flags().BoolVarP(&force, "force", "f", false, "force")
}
// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"github.com/banzaicloud/banzai-cli/internal/cli"
	"github.com/spf13/cobra"
)

// NewDeploymentCommand returns a `*cobra.Command` for `banzai cluster deployment` subcommands.
func NewDeploymentCommand(banzaiCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "deployment",
		Aliases:       []string{"deployments", "deploy"},
		Short:         "Manage deployments",
	}

	cmd.AddCommand(NewDeploymentListCommand(banzaiCli))
	return cmd
}


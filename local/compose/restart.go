/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package compose

import (
	"context"

	"github.com/docker/compose-cli/api/compose"

	"github.com/compose-spec/compose-go/types"
)

func (s *composeService) Restart(ctx context.Context, project *types.Project, options compose.RestartOptions) error {
	ctx, err := s.getUpdatedContainersStateContext(ctx, project.Name)
	if err != nil {
		return err
	}
	err = InDependencyOrder(ctx, project, func(c context.Context, service types.ServiceConfig) error {
		return s.restartService(ctx, service.Name, options.Timeout)
	})
	if err != nil {
		return err
	}
	return nil
}

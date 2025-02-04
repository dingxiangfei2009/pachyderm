package clusterstate

import (
	"context"

	"github.com/pachyderm/pachyderm/v2/src/internal/migrations"
	authserver "github.com/pachyderm/pachyderm/v2/src/server/auth/server"
)

var state_2_3_0 migrations.State = state_2_1_0.
	Apply("Add internal auth user as a cluster admin", func(ctx context.Context, env migrations.Env) error {
		return authserver.InternalAuthUserPermissions(env.Tx)
	})
	// DO NOT MODIFY THIS STATE
	// IT HAS ALREADY SHIPPED IN A RELEASE

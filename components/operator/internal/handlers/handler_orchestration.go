package handlers

import (
	stackv1beta3 "github.com/formancehq/operator/apis/stack/v1beta3"
	"github.com/formancehq/operator/internal/modules"
)

func init() {
	orchestrationEnvVars := func(resolveContext modules.ContainerResolutionContext) modules.ContainerEnv {
		return modules.NewEnv().Append(
			modules.Env("POSTGRES_DSN", "$(POSTGRES_URI)"),
			modules.Env("TEMPORAL_TASK_QUEUE", resolveContext.Stack.Name),
			modules.Env("TEMPORAL_ADDRESS", resolveContext.Configuration.Spec.Temporal.Address),
			modules.Env("TEMPORAL_NAMESPACE", resolveContext.Configuration.Spec.Temporal.Namespace),
			modules.Env("TEMPORAL_SSL_CLIENT_KEY", resolveContext.Configuration.Spec.Temporal.TLS.Key),
			modules.Env("TEMPORAL_SSL_CLIENT_CERT", resolveContext.Configuration.Spec.Temporal.TLS.CRT),
			modules.Env("STACK_CLIENT_ID", resolveContext.Stack.Status.StaticAuthClients["orchestration"].ID),
			modules.Env("STACK_CLIENT_SECRET", resolveContext.Stack.Status.StaticAuthClients["orchestration"].Secrets[0]),
		)
	}
	modules.Register("orchestration", modules.Module{
		Postgres: func(ctx modules.Context) stackv1beta3.PostgresConfig {
			return ctx.Configuration.Spec.Services.Orchestration.Postgres
		},
		Services: func(ctx modules.Context) modules.Services {
			return modules.Services{
				{
					Port: 8080,
					AuthConfiguration: func(resolveContext modules.PrepareContext) stackv1beta3.ClientConfiguration {
						return stackv1beta3.NewClientConfiguration()
					},
					ExposeHTTP:              true,
					HasVersionEndpoint:      true,
					InjectPostgresVariables: true,
					Container: func(resolveContext modules.ContainerResolutionContext) modules.Container {
						return modules.Container{
							Env:   orchestrationEnvVars(resolveContext),
							Image: modules.GetImage("orchestration", resolveContext.Versions.Spec.Orchestration),
						}
					},
				},
				{
					Name:                    "worker",
					InjectPostgresVariables: true,
					Container: func(resolveContext modules.ContainerResolutionContext) modules.Container {
						return modules.Container{
							Env:      orchestrationEnvVars(resolveContext),
							Image:    modules.GetImage("orchestration", resolveContext.Versions.Spec.Orchestration),
							Args:     []string{"worker"},
							Liveness: modules.LivenessDisable,
						}
					},
				},
			}
		},
	})
}

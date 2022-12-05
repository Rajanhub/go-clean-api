package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewRouter),
	fx.Provide(NewMigration),
	fx.Provide(NewFBApp),
	fx.Provide(NewFBAuth),
	fx.Provide(NewFCMClient),
	fx.Provide(NewFirestoreClient),
	fx.Provide(NewBucketStorage),
)

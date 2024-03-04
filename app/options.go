package app

import (
	"embed"
)

type Options struct {
	Locales       *embed.FS
	Seed          AppContextAwareFunc
	SetupApiGroup ApiGroupAwareFunc
	FinishSetup   AppContextAwareFunc
}

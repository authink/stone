package app

import (
	"embed"
)

type Options struct {
	Locales       *embed.FS
	Seed          AppContextAwareFunc
	FinishSetup   AppContextAwareFunc
	SetupAPIGroup APIGroupAwareFunc
}

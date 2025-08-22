package config

import (
	"strings"

	"github.com/MunifTanjim/stremthru/internal/util"
)

type stremioConfigTorz struct {
	LazyPull bool
}

type stremioConfigWrap struct {
	PublicMaxUpstreamCount int
	PublicMaxStoreCount    int
}

type StremioConfig struct {
	Torz stremioConfigTorz
	Wrap stremioConfigWrap
}

func parseStremio() StremioConfig {
	torzLazyPull := strings.ToLower(getEnv("STREMTHRU_STREMIO_TORZ_LAZY_PULL"))
	stremio := StremioConfig{
		Torz: stremioConfigTorz{
			LazyPull: torzLazyPull == "true",
		},
		Wrap: stremioConfigWrap{
			PublicMaxUpstreamCount: util.MustParseInt(getEnv("STREMTHRU_STREMIO_WRAP_PUBLIC_MAX_UPSTREAM_COUNT")),
			PublicMaxStoreCount:    util.MustParseInt(getEnv("STREMTHRU_STREMIO_WRAP_PUBLIC_MAX_STORE_COUNT")),
		},
	}
	return stremio
}

var Stremio = parseStremio()

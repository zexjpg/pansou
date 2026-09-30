package config

import (
	_ "embed"
	"strings"
)

// defaultChannelsEmbed 是发布版的全量频道清单，与 docker-compose.yml 的 CHANNELS 保持一致，
// 去重后为 111 个。CHANNELS 环境变量未设置时，发布版默认启用全部频道。
//
// 如需回退到上游旧的单频道默认，显式设置 CHANNELS=tgsearchers7 即可。
//
//go:embed default_channels.txt
var defaultChannelsEmbed string

// DefaultAllChannels 解析自 default_channels.txt 的全量频道列表。
var DefaultAllChannels = parseChannelLines(defaultChannelsEmbed)

func parseChannelLines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	parts := strings.Split(s, "\n")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

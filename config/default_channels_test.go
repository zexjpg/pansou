package config

import (
	"slices"
	"testing"
)

// 无 CHANNELS 环境变量时，发布版默认启用全部频道（default_channels.txt，111 个）。
// 把这个值钉住：改默认值必须同步更新 default_channels.txt 与 README。
func TestDefaultChannelsWhenEnvUnset(t *testing.T) {
	t.Setenv("CHANNELS", "")
	got := getDefaultChannels()
	if len(got) < 100 {
		t.Fatalf("无 CHANNELS 时应默认全量频道，实际仅 %d 个: %v", len(got), got)
	}
	if !slices.Contains(got, "tgsearchers7") {
		t.Errorf("默认频道列表应包含 tgsearchers7")
	}
}

func TestChannelsFromEnvAreSplit(t *testing.T) {
	t.Setenv("CHANNELS", "a,b,c")
	if got, want := getDefaultChannels(), []string{"a", "b", "c"}; !slices.Equal(got, want) {
		t.Errorf("CHANNELS 应按逗号切分，期望 %v，实际 %v", want, got)
	}
}

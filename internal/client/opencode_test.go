package client

import (
	"fmt"
	"testing"

	"oc-go-cc/internal/config"
)

func TestIsAnthropicModelOnlyRoutesNativeAnthropicModels(t *testing.T) {
	tests := []struct {
		name    string
		modelID string
		want    bool
	}{
		{
			name:    "minimax m2.5 uses anthropic endpoint",
			modelID: "minimax-m2.5",
			want:    true,
		},
		{
			name:    "minimax m2.7 uses anthropic endpoint",
			modelID: "minimax-m2.7",
			want:    true,
		},
		{
			name:    "minimax m3 uses anthropic endpoint",
			modelID: "minimax-m3",
			want:    true,
		},
		{
			name:    "deepseek pro uses openai endpoint",
			modelID: "deepseek-v4-pro",
			want:    false,
		},
		{
			name:    "deepseek flash uses openai endpoint",
			modelID: "deepseek-v4-flash",
			want:    false,
		},
		{
			name:    "kimi k2.6 uses openai endpoint",
			modelID: "kimi-k2.6",
			want:    false,
		},
		{
			name:    "glm-5.1 uses openai endpoint",
			modelID: "glm-5.1",
			want:    false,
		},
		{
			name:    "kimi-k2.5 uses openai endpoint",
			modelID: "kimi-k2.5",
			want:    false,
		},
		{
			name:    "mimo-v2.5 uses openai endpoint",
			modelID: "mimo-v2.5",
			want:    false,
		},
		{
			name:    "mimo-v2.5-pro uses openai endpoint",
			modelID: "mimo-v2.5-pro",
			want:    false,
		},
		{
			name:    "qwen3.5-plus uses anthropic endpoint",
			modelID: "qwen3.5-plus",
			want:    true,
		},
		{
			name:    "qwen3.6-plus uses anthropic endpoint",
			modelID: "qwen3.6-plus",
			want:    true,
		},
		{
			name:    "qwen3.7-plus uses anthropic endpoint",
			modelID: "qwen3.7-plus",
			want:    true,
		},
		{
			name:    "qwen3.7-max uses anthropic endpoint",
			modelID: "qwen3.7-max",
			want:    true,
		},
		{
			name:    "claude-sonnet-4-5 uses anthropic endpoint",
			modelID: "claude-sonnet-4-5",
			want:    true,
		},
		{
			name:    "claude-opus-4-7 uses anthropic endpoint",
			modelID: "claude-opus-4-7",
			want:    true,
		},
		{
			name:    "claude-haiku-4-5 uses anthropic endpoint",
			modelID: "claude-haiku-4-5",
			want:    true,
		},
		{
			name:    "claude-3-5-haiku uses anthropic endpoint",
			modelID: "claude-3-5-haiku",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnthropicModel(tt.modelID); got != tt.want {
				t.Fatalf("IsAnthropicModel(%q) = %v, want %v", tt.modelID, got, tt.want)
			}
		})
	}
}

func TestProvider(t *testing.T) {
	tests := []struct {
		name     string
		model    config.ModelConfig
		expected string
	}{
		{
			name:     "empty provider defaults to opencode-go",
			model:    config.ModelConfig{ModelID: "test-model"},
			expected: ProviderOpenCodeGo,
		},
		{
			name:     "explicit opencode-go provider",
			model:    config.ModelConfig{Provider: ProviderOpenCodeGo, ModelID: "test-model"},
			expected: ProviderOpenCodeGo,
		},
		{
			name:     "explicit opencode-zen provider",
			model:    config.ModelConfig{Provider: ProviderOpenCodeZen, ModelID: "test-model"},
			expected: ProviderOpenCodeZen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Provider(tt.model); got != tt.expected {
				t.Fatalf("Provider() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsZen(t *testing.T) {
	tests := []struct {
		name     string
		model    config.ModelConfig
		expected bool
	}{
		{
			name:     "opencode-go is not zen",
			model:    config.ModelConfig{Provider: ProviderOpenCodeGo},
			expected: false,
		},
		{
			name:     "opencode-zen is zen",
			model:    config.ModelConfig{Provider: ProviderOpenCodeZen},
			expected: true,
		},
		{
			name:     "empty provider is not zen",
			model:    config.ModelConfig{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZen(tt.model); got != tt.expected {
				t.Fatalf("IsZen() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestClassifyEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		modelID  string
		expected EndpointType
	}{
		{
			name:     "minimax m2.5 uses chat completions on Zen",
			modelID:  "minimax-m2.5",
			expected: EndpointChatCompletions,
		},
		{
			name:     "minimax m2.7 uses chat completions on Zen",
			modelID:  "minimax-m2.7",
			expected: EndpointChatCompletions,
		},
		{
			name:     "minimax m3 uses chat completions on Zen",
			modelID:  "minimax-m3",
			expected: EndpointChatCompletions,
		},
		{
			name:     "qwen3.5-plus uses anthropic endpoint",
			modelID:  "qwen3.5-plus",
			expected: EndpointAnthropic,
		},
		{
			name:     "qwen3.6-plus uses anthropic endpoint",
			modelID:  "qwen3.6-plus",
			expected: EndpointAnthropic,
		},
		{
			name:     "qwen3.7-plus uses anthropic endpoint",
			modelID:  "qwen3.7-plus",
			expected: EndpointAnthropic,
		},
		{
			name:     "qwen3.7-max uses anthropic endpoint",
			modelID:  "qwen3.7-max",
			expected: EndpointAnthropic,
		},
		{
			name:     "gemini-3.5-flash uses gemini endpoint",
			modelID:  "gemini-3.5-flash",
			expected: EndpointGemini,
		},
		{
			name:     "gemini-3.1-pro uses gemini endpoint",
			modelID:  "gemini-3.1-pro",
			expected: EndpointGemini,
		},
		{
			name:     "gemini-3-flash uses gemini endpoint",
			modelID:  "gemini-3-flash",
			expected: EndpointGemini,
		},
		{
			name:     "gpt-5.5 uses responses endpoint",
			modelID:  "gpt-5.5",
			expected: EndpointResponses,
		},
		{
			name:     "gpt-5.4 uses responses endpoint",
			modelID:  "gpt-5.4",
			expected: EndpointResponses,
		},
		{
			name:     "gpt-5 uses responses endpoint",
			modelID:  "gpt-5",
			expected: EndpointResponses,
		},
		{
			name:     "kimi-k2.6 uses chat completions endpoint",
			modelID:  "kimi-k2.6",
			expected: EndpointChatCompletions,
		},
		{
			name:     "kimi-k2.5 uses chat completions endpoint",
			modelID:  "kimi-k2.5",
			expected: EndpointChatCompletions,
		},
		{
			name:     "mimo-v2.5 uses chat completions endpoint",
			modelID:  "mimo-v2.5",
			expected: EndpointChatCompletions,
		},
		{
			name:     "mimo-v2.5-pro uses chat completions endpoint",
			modelID:  "mimo-v2.5-pro",
			expected: EndpointChatCompletions,
		},
		{
			name:     "glm-5.1 uses chat completions endpoint",
			modelID:  "glm-5.1",
			expected: EndpointChatCompletions,
		},
		{
			name:     "deepseek-v4-flash uses chat completions endpoint",
			modelID:  "deepseek-v4-flash",
			expected: EndpointChatCompletions,
		},
		{
			name:     "grok-build-0.1 uses chat completions endpoint",
			modelID:  "grok-build-0.1",
			expected: EndpointChatCompletions,
		},
		{
			name:     "big-pickle uses chat completions endpoint",
			modelID:  "big-pickle",
			expected: EndpointChatCompletions,
		},
		{
			name:     "north-mini-code-free uses chat completions endpoint",
			modelID:  "north-mini-code-free",
			expected: EndpointChatCompletions,
		},
		{
			name:     "deepseek-v4-flash-free uses chat completions endpoint",
			modelID:  "deepseek-v4-flash-free",
			expected: EndpointChatCompletions,
		},
		{
			name:     "claude-sonnet-4-5 uses anthropic endpoint",
			modelID:  "claude-sonnet-4-5",
			expected: EndpointAnthropic,
		},
		{
			name:     "claude-opus-4-7 uses anthropic endpoint",
			modelID:  "claude-opus-4-7",
			expected: EndpointAnthropic,
		},
		{
			name:     "claude-haiku-4-5 uses anthropic endpoint",
			modelID:  "claude-haiku-4-5",
			expected: EndpointAnthropic,
		},
		{
			name:     "unknown model uses chat completions endpoint",
			modelID:  "unknown-model",
			expected: EndpointChatCompletions,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClassifyEndpoint(tt.modelID); got != tt.expected {
				t.Fatalf("ClassifyEndpoint(%q) = %v, want %v", tt.modelID, got, tt.expected)
			}
		})
	}
}

func TestIsGeminiModel(t *testing.T) {
	tests := []struct {
		modelID string
		want    bool
	}{
		{"gemini-3.5-flash", true},
		{"gemini-3.1-pro", true},
		{"gemini-3-flash", true},
		{"kimi-k2.6", false},
		{"glm-5.1", false},
		{"gpt-5.5", false},
	}

	for _, tt := range tests {
		t.Run(tt.modelID, func(t *testing.T) {
			if got := isGeminiModel(tt.modelID); got != tt.want {
				t.Fatalf("isGeminiModel(%q) = %v, want %v", tt.modelID, got, tt.want)
			}
		})
	}
}

func TestIsResponsesModel(t *testing.T) {
	tests := []struct {
		modelID string
		want    bool
	}{
		{"gpt-5.5", true},
		{"gpt-5.5-pro", true},
		{"gpt-5.4", true},
		{"gpt-5.4-pro", true},
		{"gpt-5.4-mini", true},
		{"gpt-5.4-nano", true},
		{"gpt-5.3-codex", true},
		{"gpt-5.3-codex-spark", true},
		{"gpt-5.2", true},
		{"gpt-5.2-codex", true},
		{"gpt-5.1", true},
		{"gpt-5.1-codex", true},
		{"gpt-5.1-codex-max", true},
		{"gpt-5.1-codex-mini", true},
		{"gpt-5", true},
		{"gpt-5-codex", true},
		{"gpt-5-nano", true},
		{"kimi-k2.6", false},
		{"glm-5.1", false},
		{"gemini-3.5-flash", false},
	}

	for _, tt := range tests {
		t.Run(tt.modelID, func(t *testing.T) {
			if got := isResponsesModel(tt.modelID); got != tt.want {
				t.Fatalf("isResponsesModel(%q) = %v, want %v", tt.modelID, got, tt.want)
			}
		})
	}
}

func TestNextAPIKey_RoundRobin(t *testing.T) {
	cfg := &config.Config{
		APIKeys: []string{"key-a", "key-b", "key-c"},
	}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{
		atomic: atomicCfg,
	}

	// With 3 keys, iteration 0..5 should cycle key-a, key-b, key-c, key-a, key-b, key-c
	expected := []string{"key-a", "key-b", "key-c", "key-a", "key-b", "key-c"}
	for i, want := range expected {
		if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), ""); got != want {
			t.Errorf("iteration %d: nextAPIKey() = %q, want %q", i, got, want)
		}
	}
}

func TestNextAPIKey_SingleKey(t *testing.T) {
	cfg := &config.Config{APIKey: "single"}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	for i := 0; i < 5; i++ {
		if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), ""); got != "single" {
			t.Errorf("iteration %d: nextAPIKey() = %q, want %q", i, got, "single")
		}
	}
}

func TestNextAPIKey_EmptyKeys(t *testing.T) {
	cfg := &config.Config{APIKey: ""}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), ""); got != "" {
		t.Errorf("nextAPIKey() = %q, want empty string", got)
	}
}

func TestNextAPIKey_ConcurrentSafety(t *testing.T) {
	cfg := &config.Config{
		APIKeys: []string{"k1", "k2", "k3"},
	}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	const goroutines = 3
	const callsPerGoroutine = 100
	results := make(chan string, goroutines*callsPerGoroutine)

	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < callsPerGoroutine; i++ {
				results <- c.nextAPIKey(cfg.EffectiveAPIKeys(), "")
			}
		}()
	}

	seen := make(map[string]int)
	for i := 0; i < goroutines*callsPerGoroutine; i++ {
		key := <-results
		seen[key]++
	}

	// Each key should be seen exactly (goroutines*callsPerGoroutine)/3 times
	total := goroutines * callsPerGoroutine
	expectedPerKey := total / len(cfg.APIKeys)
	for _, key := range cfg.APIKeys {
		if seen[key] != expectedPerKey {
			t.Errorf("key %q seen %d times, want %d", key, seen[key], expectedPerKey)
		}
	}
}

// TestNextAPIKey_StickyDeterminism verifies that a non-empty sticky key
// always resolves to the same upstream key across many invocations.
func TestNextAPIKey_StickyDeterminism(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{"key-a", "key-b", "key-c"}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	first := c.nextAPIKey(cfg.EffectiveAPIKeys(), "token-X")
	for i := 0; i < 100; i++ {
		if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), "token-X"); got != first {
			t.Errorf("iteration %d: sticky nextAPIKey drifted: got %q, want %q", i, got, first)
		}
	}
}

// TestNextAPIKey_StickyDistribution verifies that a 3-key pool, fed many
// distinct tokens, hits more than one key (so hashing is not constant) and
// each token's key matches the FNV-1a index helper.
func TestNextAPIKey_StickyDistribution(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{"key-a", "key-b", "key-c"}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	seen := make(map[string]struct{})
	for i := 0; i < 32; i++ {
		token := fmt.Sprintf("token-%d", i)
		want := cfg.APIKeys[stickyKeyIndex(token, uint64(len(cfg.APIKeys)))]
		got := c.nextAPIKey(cfg.EffectiveAPIKeys(), token)
		if got != want {
			t.Errorf("token %q: got %q, want %q (index helper mismatch)", token, got, want)
		}
		seen[got] = struct{}{}
	}
	if len(seen) < 2 {
		t.Errorf("expected sticky hashing to map 32 tokens onto at least 2 keys; saw %d", len(seen))
	}
}

// TestNextAPIKey_StickyEmptyFallsBackToRoundRobin verifies that an empty
// sticky key preserves the original round-robin cycle (backward compatibility
// for callers that haven't opted into sticky routing).
func TestNextAPIKey_StickyEmptyFallsBackToRoundRobin(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{"key-a", "key-b", "key-c"}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	expected := []string{"key-a", "key-b", "key-c", "key-a", "key-b", "key-c"}
	for i, want := range expected {
		if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), ""); got != want {
			t.Errorf("iteration %d: got %q, want %q", i, got, want)
		}
	}
}

// TestNextAPIKey_StickySingleKeyPool verifies that a 1-key pool always
// returns that key regardless of the sticky token.
func TestNextAPIKey_StickySingleKeyPool(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{"only"}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	for _, token := range []string{"", "alpha", "beta", "anything"} {
		if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), token); got != "only" {
			t.Errorf("token %q: got %q, want %q", token, got, "only")
		}
	}
}

// TestNextAPIKey_StickyNoInteractionWithCounter verifies that interleaving
// sticky and round-robin calls in one client does not break either path:
// sticky calls are deterministic, "" calls still advance the counter.
func TestNextAPIKey_StickyNoInteractionWithCounter(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{"k1", "k2", "k3"}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	// Two sticky calls first — must pick the same key for "alpha".
	const token = "alpha"
	want := cfg.APIKeys[stickyKeyIndex(token, uint64(len(cfg.APIKeys)))]
	if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), token); got != want {
		t.Fatalf("alpha first: got %q, want %q", got, want)
	}
	if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), token); got != want {
		t.Fatalf("alpha second: got %q, want %q", got, want)
	}
	// Now drain the counter with three "" calls. Order is implementation-
	// defined but each key must appear exactly once.
	got := []string{
		c.nextAPIKey(cfg.EffectiveAPIKeys(), ""),
		c.nextAPIKey(cfg.EffectiveAPIKeys(), ""),
		c.nextAPIKey(cfg.EffectiveAPIKeys(), ""),
	}
	seen := make(map[string]bool)
	for _, k := range got {
		if seen[k] {
			t.Errorf("counter walk produced duplicate key %q in %v", k, got)
		}
		seen[k] = true
	}
	if len(seen) != 3 {
		t.Errorf("expected 3 distinct keys in counter walk, got %v", got)
	}
	// Sticky call afterwards is still deterministic.
	if got := c.nextAPIKey(cfg.EffectiveAPIKeys(), token); got != want {
		t.Errorf("alpha post-counter: got %q, want %q", got, want)
	}
}

// TestResolveAPIKey_ExactMappingWins verifies that an explicit token->index
// mapping in config is honored verbatim, regardless of what the FNV-1a hash
// would have produced for the same token.
func TestResolveAPIKey_ExactMappingWins(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{"sk-alice", "sk-bob", "sk-carol"}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	mappings := map[string]int{
		"customerA-token": 0, // -> sk-alice
		"customerB-token": 1, // -> sk-bob
		"customerC-token": 2, // -> sk-carol
	}

	cases := []struct {
		token string
		want  string
	}{
		{"customerA-token", "sk-alice"},
		{"customerB-token", "sk-bob"},
		{"customerC-token", "sk-carol"},
	}
	for _, tc := range cases {
		if got := c.ResolveAPIKey(cfg.EffectiveAPIKeys(), tc.token, mappings); got != tc.want {
			t.Errorf("token %q: got %q, want %q (explicit mapping should win)", tc.token, got, tc.want)
		}
	}
}

// TestResolveAPIKey_MappingMissFallsBackToHash verifies that an unmapped
// token still gets stable bucketing via FNV-1a.
func TestResolveAPIKey_MappingMissFallsBackToHash(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{"sk-alice", "sk-bob", "sk-carol"}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	mappings := map[string]int{"customerA-token": 0}

	// Unmapped token "customerZ-token" must equal what nextAPIKey would
	// have produced with the same input — i.e., deterministic hash bucket.
	want := c.nextAPIKey(cfg.EffectiveAPIKeys(), "customerZ-token")
	if got := c.ResolveAPIKey(cfg.EffectiveAPIKeys(), "customerZ-token", mappings); got != want {
		t.Errorf("unmapped token: got %q, want %q (hash fallback)", got, want)
	}
	// And it must be stable across calls.
	for i := 0; i < 5; i++ {
		if got := c.ResolveAPIKey(cfg.EffectiveAPIKeys(), "customerZ-token", mappings); got != want {
			t.Errorf("iteration %d: hash fallback drifted, got %q, want %q", i, got, want)
		}
	}
}

// TestResolveAPIKey_OutOfRangeMappingFallsThrough verifies that a misconfigured
// mapping (index out of range) is treated as a miss and falls through to the
// hash path, so a bad config cannot panic and cannot strand a token to "".
func TestResolveAPIKey_OutOfRangeMappingFallsThrough(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{"sk-alice", "sk-bob"}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	mappings := map[string]int{
		"overshoot": 99,  // way past end
		"negative":  -1,  // negative
		"edge":      2,   // exactly len(keys) (out of range)
	}
	for _, tok := range []string{"overshoot", "negative", "edge"} {
		want := c.nextAPIKey(cfg.EffectiveAPIKeys(), tok)
		got := c.ResolveAPIKey(cfg.EffectiveAPIKeys(), tok, mappings)
		if got != want {
			t.Errorf("token %q with out-of-range mapping: got %q, want hash fallback %q", tok, got, want)
		}
		if got == "" {
			t.Errorf("token %q: out-of-range mapping must not produce empty key", tok)
		}
	}
}

// TestResolveAPIKey_EmptyPoolAndEmptyToken verifies edge cases:
// empty key pool -> "", empty token -> round-robin.
func TestResolveAPIKey_EmptyPoolAndEmptyToken(t *testing.T) {
	cfg := &config.Config{APIKeys: []string{}}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	if got := c.ResolveAPIKey(cfg.EffectiveAPIKeys(), "any-token", map[string]int{"any-token": 0}); got != "" {
		t.Errorf("empty pool: got %q, want \"\"", got)
	}

	cfg2 := &config.Config{APIKeys: []string{"k1", "k2", "k3"}}
	atomicCfg2 := config.NewAtomicConfig(cfg2, "")
	c2 := &OpenCodeClient{atomic: atomicCfg2}
	// Empty token falls back to nextAPIKey which is round-robin.
	got1 := c2.ResolveAPIKey(cfg2.EffectiveAPIKeys(), "", nil)
	got2 := c2.ResolveAPIKey(cfg2.EffectiveAPIKeys(), "", nil)
	if got1 != "k1" || got2 != "k2" {
		t.Errorf("empty token: got (%q, %q), want (k1, k2) — round-robin", got1, got2)
	}
}

// goSubscriptionModels is the authoritative list of model IDs included in the
// OpenCode Go subscription, paired with the endpoint each one uses per the docs
// (https://opencode.ai/docs/go). anthropic=true means the model is served on
// the Go Anthropic endpoint (/v1/messages); false means the Go chat-completions
// endpoint (/v1/chat/completions). This table is the routing contract for Go
// models — keep it in sync with the docs when OpenCode adds models.
var goSubscriptionModels = []struct {
	modelID   string
	anthropic bool
}{
	// OpenAI-compatible endpoint: https://opencode.ai/zen/go/v1/chat/completions
	{"glm-5.1", false},
	{"glm-5", false},
	{"kimi-k2.7", false},
	{"kimi-k2.6", false},
	{"deepseek-v4-pro", false},
	{"deepseek-v4-flash", false},
	{"mimo-v2.5", false},
	{"mimo-v2.5-pro", false},
	// Anthropic endpoint: https://opencode.ai/zen/go/v1/messages
	{"minimax-m3", true},
	{"minimax-m2.7", true},
	{"minimax-m2.5", true},
	{"qwen3.7-max", true},
	{"qwen3.7-plus", true},
	{"qwen3.6-plus", true},
}

// TestGoSubscriptionModelsRouteToGoEndpoint tests every model in the OpenCode
// Go subscription individually, asserting for each that:
//   - IsAnthropicModel classifies it to the endpoint the docs specify, and
//   - an opencode-go ModelConfig resolves (via getEndpoint) to the matching Go
//     base URL — and is never treated as a Zen model or routed to a Zen URL.
//
// Guards the regression where a Go model's model_overrides entry pointed at
// provider=opencode-zen (whose account had no balance), causing a 401 and a
// silent fallback to an unrelated model.
func TestGoSubscriptionModelsRouteToGoEndpoint(t *testing.T) {
	const (
		goBaseURL          = "https://opencode.ai/zen/go/v1/chat/completions"
		goAnthropicBaseURL = "https://opencode.ai/zen/go/v1/messages"
		zenBaseURL         = "https://opencode.ai/zen/v1/chat/completions"
		zenAnthropicURL    = "https://opencode.ai/zen/v1/messages"
	)

	cfg := &config.Config{
		OpenCodeGo: config.OpenCodeGoConfig{
			BaseURL:          goBaseURL,
			AnthropicBaseURL: goAnthropicBaseURL,
		},
		OpenCodeZen: config.OpenCodeZenConfig{
			BaseURL:          zenBaseURL,
			AnthropicBaseURL: zenAnthropicURL,
		},
		APIKeys: []string{"go-key"},
	}
	atomicCfg := config.NewAtomicConfig(cfg, "")
	c := &OpenCodeClient{atomic: atomicCfg}

	for _, m := range goSubscriptionModels {
		t.Run(m.modelID, func(t *testing.T) {
			// Endpoint classification matches the docs.
			if got := IsAnthropicModel(m.modelID); got != m.anthropic {
				t.Errorf("IsAnthropicModel(%q) = %v, want %v", m.modelID, got, m.anthropic)
			}

			modelCfg := config.ModelConfig{Provider: ProviderOpenCodeGo, ModelID: m.modelID}

			// A Go-provider config must not be treated as Zen.
			if IsZen(modelCfg) {
				t.Errorf("IsZen(%q) = true; Go subscription model must not be Zen", m.modelID)
			}

			// And it must resolve to the matching Go base URL.
			ep := c.getEndpoint(m.modelID, modelCfg, "")
			wantBase := goBaseURL
			if m.anthropic {
				wantBase = goAnthropicBaseURL
			}
			if ep.BaseURL != wantBase {
				t.Errorf("getEndpoint(%q) BaseURL = %q, want %q", m.modelID, ep.BaseURL, wantBase)
			}
			if ep.APIKey != "go-key" {
				t.Errorf("getEndpoint(%q) APIKey = %q, want %q", m.modelID, ep.APIKey, "go-key")
			}
		})
	}
}

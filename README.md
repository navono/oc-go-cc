# oc-go-cc

A Go CLI proxy that lets you use your [OpenCode Go](https://opencode.ai/docs/go/) or [OpenCode Zen](https://opencode.ai/docs/zen/) subscription with [Claude Code](https://docs.anthropic.com/en/docs/claude-code).

`oc-go-cc` sits between Claude Code and OpenCode, intercepting Anthropic API requests, transforming them to the appropriate format (OpenAI, Responses, or Gemini), and forwarding them to OpenCode's endpoints. Claude Code thinks it's talking to Anthropic — but your requests go to affordable open models instead.

## Why?

OpenCode Go gives you access to powerful open coding models for **$5/month** (then $10/month). OpenCode Zen provides curated, tested models with pay-as-you-go pricing. This proxy makes both work seamlessly with Claude Code's interface — no patches, no forks, just set two environment variables and go.

## Features

- **Transparent Proxy** — Claude Code sends Anthropic-format requests, proxy transforms to OpenAI/Responses/Gemini format and back
- **Dual Provider Support** — Route models through OpenCode Go or OpenCode Zen based on your needs
- **Model Routing** — Automatically routes to different models based on context (default, thinking, long context, background)
- **Fallback Chains** — If a model fails, automatically tries the next one in your configured chain
- **Circuit Breaker** — Tracks model health and skips failing models to avoid latency spikes
- **Real-time Streaming** — Full SSE streaming with live format transformation
- **Tool Calling** — Proper Anthropic tool_use/tool_result <-> OpenAI/Gemini function calling translation
- **Token Counting** — Uses tiktoken (cl100k_base) for accurate token counting and context threshold detection
- **JSON Configuration** — Flexible config file with environment variable overrides and `${VAR}` interpolation
- **Hot Reload** — Watch config file for changes and reload automatically (off by default)
- **Background Mode** — Run as daemon detached from terminal
- **Auto-start on Login** — Launch on system startup via launchd (macOS)

## Quick Start

### 1. Install

```bash
# macOS / Linux
brew tap samueltuyizere/tap && brew install oc-go-cc

# Windows
scoop bucket add oc-go-cc https://github.com/samueltuyizere/scoop-bucket && scoop install oc-go-cc

# Docker (with Makefile)
cp .env.example .env                    # then put your API key in .env
make docker-up

# Docker (manual)
cp .env.example .env
docker build -t oc-go-cc .
docker run -d --restart unless-stopped --name oc-go-cc --env-file .env -p 3456:3456 oc-go-cc
```

Or see [INSTALLATION.md](INSTALLATION.md) for more options.

### 2. Initialize Configuration

```bash
oc-go-cc init
```

Creates a default config at `~/.config/oc-go-cc/config.json`. Edit it to add your API key, or set the environment variable:

```bash
export OC_GO_CC_API_KEY=sk-opencode-your-key-here
```

### 3. Start the Proxy

```bash
oc-go-cc serve
```

Stop the Docker container (if using Docker):

```bash
make docker-stop
```

### 4. Configure Claude Code

```bash
export ANTHROPIC_BASE_URL=http://127.0.0.1:3456
export ANTHROPIC_AUTH_TOKEN=unused
```

## Multiple API Keys & Sticky Routing

By default `oc-go-cc` rotates through a single `api_key`. To distribute load
across multiple OpenCode accounts (e.g. per-customer billing or per-team
quotas), define an `api_keys` array. Set `"sticky_key_enabled": true` if you
want the same inbound client to always land on the same upstream key — the
inbound `ANTHROPIC_AUTH_TOKEN` (or the `x-api-key` / `Authorization: Bearer`
header that Claude Code sends) is hashed (FNV-1a 32-bit) to pick the key
deterministically, and the choice is locked for the entire fallback chain so
a request that fails over from primary to backup model stays on the same
upstream account.

```json
{
  "api_keys": [
    "${OC_GO_CC_API_KEY_1}",
    "${OC_GO_CC_API_KEY_2}",
    "${OC_GO_CC_API_KEY_3}"
  ],
  "sticky_key_enabled": true
}
```

- `api_keys` takes precedence over `api_key`; omit `api_key` when using the
  array form.
- When `sticky_key_enabled` is `false` (the default) the keys are picked by
  round-robin — the same as before this feature.
- When `sticky_key_enabled` is `true` but the inbound request has no auth
  token, the proxy falls back to round-robin so requests are never dropped.

### Pin a specific `ANTHROPIC_AUTH_TOKEN` to a specific key

For per-customer / per-team billing or quota isolation, add an explicit
`sticky_key_mappings` table. The proxy checks this map first — an exact
match wins regardless of the FNV-1a hash result. Unmapped tokens still
fall back to hash bucketing.

```json
{
  "api_keys": [
    "sk-alice-account",
    "sk-bob-account",
    "sk-carol-account"
  ],
  "sticky_key_enabled": true,
  "sticky_key_mappings": {
    "customerA-token": 0,
    "customerB-token": 1,
    "customerC-token": 2
  }
}
```

The numbers are **indices into `api_keys`** (0-based). The token strings are
the exact `ANTHROPIC_AUTH_TOKEN` values your clients send. With this config
in place, every request carrying `customerA-token` is pinned to
`sk-alice-account` for its entire fallback chain; clients that mistype or
send an unknown token are still handled via the hash bucket — they are
never dropped, just routed to a less predictable upstream key.

Misconfigured entries (out-of-range index) are silently treated as misses,
so a typo in the config cannot break the proxy.

### 5. Run Claude Code

```bash
claude
```

## CLI Commands

```
oc-go-cc serve              Start the proxy server
oc-go-cc serve -b           Start in background (detached from terminal)
oc-go-cc serve --port 8080  Start on a custom port
oc-go-cc stop               Stop the running proxy server
oc-go-cc status             Check if the proxy is running
oc-go-cc init               Create default configuration file
oc-go-cc validate           Validate configuration file
oc-go-cc models             List all available models (Go + Zen)
oc-go-cc autostart enable   Enable auto-start on login
oc-go-cc autostart disable  Disable auto-start on login
oc-go-cc autostart status   Check autostart status
oc-go-cc --version          Show version
```

## Documentation

| Document | Description |
| -------- | ----------- |
| [INSTALLATION.md](INSTALLATION.md) | Homebrew, Scoop, build from source, release binaries |
| [CONFIGURATION.md](CONFIGURATION.md) | Config file reference, env vars, model routing, fallback chains |
| [MODELS.md](MODELS.md) | Model capabilities, costs, and routing recommendations |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Development setup, architecture, how it works |
| [TROUBLESHOOTING.md](TROUBLESHOOTING.md) | Common issues and debug mode |

## License

[AGPL-3.0](LICENSE)

# OpenCode Models Guide

Comprehensive guide to OpenCode Go and Zen models with capabilities, costs, and routing recommendations.

**Sources:** [OpenCode Go Documentation](https://opencode.ai/docs/go/) | [OpenCode Zen Documentation](https://opencode.ai/docs/zen/)

## Quick Cost Comparison

> 💰 **Cost-conscious routing matters!** DeepSeek V4 Flash gives you 31,650 requests per $12, while GLM-5.1 gives you only 880 — that's **36x fewer requests** for the same budget.

| Model            | Provider      | Requests per $12 (5hr) | Cost Efficiency | Quality |
| ---------------- | ------------- | ---------------------- | --------------- | ------- |
| **Qwen3.5 Plus** | Go            | **10,200**             | ★★★★★           | ★★☆☆☆   |
| **MiniMax M2.5** | Go            | **6,300**              | ★★★★★           | ★★☆☆☆   |
| **MiniMax M2.7** | Go            | **3,400**              | ★★★★☆           | ★★★☆☆   |
| **Qwen3.6 Plus** | Go            | **3,300**              | ★★★★☆           | ★★★☆☆   |
| **MiMo-V2.5**    | Go            | **2,150**              | ★★★☆☆           | ★★★☆☆   |
| **Kimi K2.5**    | Go            | **1,850**              | ★★☆☆☆           | ★★★★☆   |
| **MiMo-V2.5-Pro**| Go            | **1,290**              | ★★☆☆☆           | ★★★★☆   |
| **Kimi K2.6**    | Go            | **~1,150**             | ★☆☆☆☆           | ★★★★★   |
| **GLM-5**        | Go            | **1,150**              | ★☆☆☆☆           | ★★★★☆   |
| **GLM-5.1**      | Go            | **880**                | ☆☆☆☆☆           | ★★★★★   |

## Providers

### OpenCode Go (`opencode-go`)

- Subscription-based ($5 first month, then $10/month)
- OpenAI Chat Completions and Anthropic Messages endpoints
- Best for: Most use cases, cost-effective models

### OpenCode Zen (`opencode-zen`)

- Pay-as-you-go pricing
- Additional endpoint formats: Responses (GPT), Gemini
- Best for: GPT models, Gemini models, premium Anthropic models

## Important: API Endpoints

⚠️ **Critical:** Not all models use the same API endpoint! oc-go-cc handles this automatically, but you should know:

### OpenCode Go Endpoints

| Models                                                                                                             | Endpoint                                         | Format                   |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------ | ------------------------ |
| GLM-5, GLM-5.1, Kimi K2.5, Kimi K2.6, MiMo-V2.5, MiMo-V2.5-Pro, DeepSeek V4 Pro, DeepSeek V4 Flash | `https://opencode.ai/zen/go/v1/chat/completions` | OpenAI-compatible        |
| **MiniMax M2.5, MiniMax M2.7, MiniMax M3, Qwen3.5 Plus, Qwen3.6 Plus, Qwen3.7 Plus, Qwen3.7 Max**          | `https://opencode.ai/zen/go/v1/messages`         | **Anthropic-compatible** |

### OpenCode Zen Endpoints

| Models                                                                           | Endpoint                                     | Format                   |
| -------------------------------------------------------------------------------- | -------------------------------------------- | ------------------------ |
| MiniMax M2.5, MiniMax M2.7, MiniMax M3, GLM-5, GLM-5.1, Kimi K2.5, Kimi K2.6, DeepSeek V4 Pro, DeepSeek V4 Flash, DeepSeek V4 Flash Free, Grok Build 0.1, Big Pickle, MiMo-V2.5 Free, North Mini Code Free, Nemotron 3 Ultra Free | `https://opencode.ai/zen/v1/chat/completions` | OpenAI-compatible        |
| **Claude models** (claude-fable-5, claude-opus-4-8, claude-opus-4-7, claude-opus-4-6, claude-opus-4-5, claude-sonnet-4-6, claude-sonnet-4-5, claude-haiku-4-5, etc.), **Qwen models** (qwen3.5-plus, qwen3.6-plus, qwen3.7-plus, qwen3.7-max) | `https://opencode.ai/zen/v1/messages`        | **Anthropic-compatible** |
| **GPT models** (gpt-5.5, gpt-5.5-pro, gpt-5.4, gpt-5.4-pro, gpt-5.4-mini, gpt-5.4-nano, gpt-5.3-codex, gpt-5.3-codex-spark, gpt-5.2, gpt-5.2-codex, gpt-5.1, gpt-5.1-codex, gpt-5.1-codex-max, gpt-5.1-codex-mini, gpt-5, gpt-5-codex, gpt-5-nano) | `https://opencode.ai/zen/v1/responses`       | **OpenAI Responses**     |
| **Gemini models** (gemini-3.5-flash, gemini-3.1-pro, gemini-3-flash)             | `https://opencode.ai/zen/v1/models/{id}`     | **Google Gemini**        |

**Why this matters:** On the Go provider, MiniMax and Qwen models use Anthropic format natively. On Zen, only Claude and Qwen use the Anthropic endpoint — MiniMax uses chat completions. oc-go-cc handles all routing automatically.

## Using OpenCode Zen

To use Zen models, set `"provider": "opencode-zen"` in your model config:

```json
{
  "models": {
    "default": {
      "provider": "opencode-zen",
      "model_id": "kimi-k2.6",
      "temperature": 0.7,
      "max_tokens": 4096
    }
  }
}
```

### Zen-Specific Models (49 total)

All OpenCode Go models are also available on Zen. Zen additionally offers:

- **Claude Models (Anthropic endpoint):** claude-fable-5, claude-opus-4-8, claude-opus-4-7, claude-opus-4-6, claude-opus-4-5, claude-opus-4-1, claude-sonnet-4-6, claude-sonnet-4-5, claude-sonnet-4, claude-haiku-4-5, claude-3-5-haiku
- **GPT Models (Responses endpoint):** gpt-5.5, gpt-5.5-pro, gpt-5.4, gpt-5.4-pro, gpt-5.4-mini, gpt-5.4-nano, gpt-5.3-codex, gpt-5.3-codex-spark, gpt-5.2, gpt-5.2-codex, gpt-5.1, gpt-5.1-codex, gpt-5.1-codex-max, gpt-5.1-codex-mini, gpt-5, gpt-5-codex, gpt-5-nano
- **Gemini Models (Gemini endpoint):** gemini-3.5-flash, gemini-3.1-pro, gemini-3-flash
- **Free Tier (chat completions):** deepseek-v4-pro, deepseek-v4-flash-free, grok-build-0.1, big-pickle, mimo-v2.5-free, north-mini-code-free, nemotron-3-ultra-free

DeepSeek V4 Pro and Flash are OpenAI-compatible on both Go and Zen providers. On Zen, DeepSeek V4 Pro is available as a free-tier model. oc-go-cc transforms Claude Code's Anthropic request into OpenAI Chat Completions format, including tools, tool results, thinking history, `reasoning_effort`, and `thinking`.

For Claude Code and OpenCode-style agent workflows, DeepSeek V4 supports max thinking mode with:

```json
{
  "model_id": "deepseek-v4-pro",
  "reasoning_effort": "max",
  "thinking": {
    "type": "enabled"
  }
}
```

Use `deepseek-v4-pro` for default, complex, thinking, and long-context routing. Use `deepseek-v4-flash` for fast, background, or subagent-style workloads.

To route DeepSeek V4 Pro through Zen (free tier) instead of Go (paid), add a `model_overrides` entry:

```json
{
  "model_overrides": {
    "deepseek-v4-pro": {
      "provider": "opencode-zen",
      "model_id": "deepseek-v4-pro",
      "temperature": 0.7,
      "max_tokens": 8192,
      "reasoning_effort": "max",
      "thinking": {
        "type": "enabled"
      }
    }
  }
}
```

## Cost-Conscious Routing Strategy

### Default to Cheap, Upgrade When Necessary

**Most requests should use cheap models.** Only upgrade to expensive models when:

1. **Task complexity demands it** (multi-step reasoning, architecture)
2. **You've tried cheaper models and they failed**
3. **Code quality is critical** (production code review)

### Recommended Routing

```json
{
  "models": {
    "background": {
      // Simple operations
      "model_id": "qwen3.7-plus",
      "max_tokens": 2048
    },
    "default": {
      // Better quality, moderate cost
      "model_id": "kimi-k2.6",
      "max_tokens": 4096
    },
    "long_context": {
      // Large files only
      "model_id": "minimax-m3",
      "context_threshold": 80000
    },
    "think": {
      // Reasoning tasks
      "model_id": "glm-5",
      "max_tokens": 8192
    },
    "complex": {
      // Complex architecture only
      "model_id": "glm-5.1",
      "max_tokens": 4096
    },
    "fast": {
      // Streaming requests (prioritize TTFT)
      "model_id": "qwen3.6-plus",
      "max_tokens": 4096
    }
  }
}
```

### Decision Tree

```
Is context > 80K tokens?
├── YES → Use MiniMax M3 (1M context, 3,200 req/$12)
│
Is it a complex task (architecture, refactoring, tool operations)?
├── YES → Use GLM-5.1 (880 req/$12)
│
Is it a reasoning/planning task?
├── YES → Use GLM-5 (1,150 req/$12)
│
Is it a simple background task (read file, grep, list dir, no tools)?
├── YES → Use Qwen3.7 Plus (4,300 req/$12)
│
Default → Use Kimi K2.6 (~1,150 req/$12, ★★★★★) or Qwen3.6 Plus (3,300 req/$12)
```

## Detailed Model Profiles

### Budget Champions 💰

#### DeepSeek V4 Flash — Fastest & Cheapest

- **Model ID:** `deepseek-v4-flash`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Cost:** **31,650 requests per $12** (best value!)
- **Context:** ~1M tokens
- **Quality:** ★★★★☆
- **Best For:**
  - Fast routing / streaming
  - Background tasks
  - Subagent-style work
  - Fallback for any scenario
- **When to Use:** When you need maximum throughput at minimum cost

#### MiMo-V2.5 — Ultra-Cheap General Purpose

- **Model ID:** `mimo-v2.5`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Cost:** **30,100 requests per $12**
- **Context:** ~128K tokens
- **Quality:** ★★★☆☆
- **Best For:**
  - Bulk operations
  - Simple coding tasks
  - Background tasks

#### Qwen3.7 Plus — Cost-Effective Workhorse

- **Model ID:** `qwen3.7-plus`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages`)
- **Cost:** **4,300 requests per $12**
- **Context:** ~256K tokens
- **Quality:** ★★★☆☆
- **Speed:** Fast
- **Best For:**
  - File reading operations
  - General coding tasks
  - Background tasks
  - Default for cost-conscious users

#### MiniMax M2.5 — Long Context on a Budget

- **Model ID:** `minimax-m2.5`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages` on Go), **OpenAI-compatible** (`/chat/completions` on Zen)
- **Cost:** **6,300 requests per $12**
- **Context:** **~1M tokens** (1 million!)
- **Quality:** ★★☆☆☆ (acceptable)
- **Speed:** Fast
- **Best For:**
  - Very large files
  - Long conversations
  - Multi-file context
- **When to Use:** When you need 1M context but want to minimize cost
- **Note:** Uses Anthropic endpoint on Go but chat completions on Zen - oc-go-cc handles this automatically

#### MiniMax M3 — Latest MiniMax, 1M Context

- **Model ID:** `minimax-m3`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages` on Go), **OpenAI-compatible** (`/chat/completions` on Zen)
- **Context:** **~1M tokens**
- **Quality:** ★★★☆☆
- **Best For:**
  - Long-context tasks requiring better quality than M2.5
  - Large codebase analysis
  - Document processing
- **When to Use:** When you need 1M context and want better quality than M2.5

### Balanced Models (Quality + Cost)

#### DeepSeek V4 Pro — Agentic Coding + Max Thinking

- **Model ID:** `deepseek-v4-pro`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Context:** **~1M tokens**
- **Quality:** ★★★★★
- **Providers:** Go (paid) or Zen (free tier)
- **Best For:**
  - Claude Code agent workflows
  - Complex implementation and debugging
  - Architecture and refactoring
  - Long-context coding tasks
  - Max thinking mode
- **Recommended Config (Go):**

  ```json
  {
    "provider": "opencode-go",
    "model_id": "deepseek-v4-pro",
    "temperature": 0.1,
    "max_tokens": 8192,
    "reasoning_effort": "max",
    "thinking": {
      "type": "enabled"
    }
  }
  ```

- **Recommended Config (Zen free tier):**

  ```json
  {
    "provider": "opencode-zen",
    "model_id": "deepseek-v4-pro",
    "temperature": 0.1,
    "max_tokens": 8192,
    "reasoning_effort": "max",
    "thinking": {
      "type": "enabled"
    }
  }
  ```

#### DeepSeek V4 Flash — Fast Agent Workloads

- **Model ID:** `deepseek-v4-flash`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Context:** **~1M tokens**
- **Quality:** ★★★★☆
- **Best For:**
  - Fast routing
  - Background tasks
  - Subagent-style work
  - Fallback for DeepSeek V4 Pro
- **Recommended Config:**

  ```json
  {
    "provider": "opencode-go",
    "model_id": "deepseek-v4-flash",
    "temperature": 0.1,
    "max_tokens": 4096,
    "reasoning_effort": "max",
    "thinking": {
      "type": "enabled"
    }
  }
  ```

#### Qwen3.6 Plus — Cost-Effective General Coding ⭐ RECOMMENDED DEFAULT

- **Model ID:** `qwen3.6-plus`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages` — Go), **Anthropic-compatible** (`/v1/messages` — Zen)
- **Cost:** **3,300 requests per $12** (3.8x more than GLM-5.1!)
- **Context:** ~128K tokens
- **Quality:** ★★★☆☆ (good enough for most tasks)
- **Speed:** Fast
- **Best For:**
  - General coding
  - Feature implementation
  - Bug fixes
  - Streaming (fast TTFT)
- **When to Use:** Fast scenario and streaming workloads

#### Qwen3.7 Max — Premium Qwen

- **Model ID:** `qwen3.7-max`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages`)
- **Cost:** **950 requests per $12**
- **Context:** ~256K tokens
- **Quality:** ★★★★★
- **Best For:**
  - Complex coding requiring top quality
  - Architecture decisions
  - Code review
- **When to Use:** When you want top Qwen quality and are willing to pay more

#### Qwen3.7 Plus — Upgraded General Coding

- **Model ID:** `qwen3.7-plus`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages`)
- **Context:** ~128K tokens
- **Quality:** ★★★★☆
- **Speed:** Fast
- **Best For:**
  - General coding with better quality than Qwen3.6
  - Feature implementation
  - Bug fixes
- **When to Use:** When you want better quality than Qwen3.6 at similar speed

#### Qwen3.7 Max — Maximum Quality Qwen

- **Model ID:** `qwen3.7-max`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages`)
- **Context:** ~128K tokens
- **Quality:** ★★★★☆
- **Best For:**
  - Complex coding tasks
  - When Qwen3.7 Plus isn't enough
- **When to Use:** When you need Qwen's best quality

#### Kimi K2.6 — Best Quality at Balanced Cost

- **Model ID:** `kimi-k2.6`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Cost:** **~1,150 requests per $12**
- **Context:** ~256K tokens
- **Quality:** ★★★★★ (excellent)
- **Speed:** Fast
- **Best For:**
  - Complex coding tasks
  - Code review
  - Architecture discussions
  - General-purpose default (best quality-to-cost ratio)
- **When to Use:** Default choice — best all-round quality

#### Kimi K2.5 — Quality + Reasonable Cost

- **Model ID:** `kimi-k2.5`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Cost:** **1,850 requests per $12**
- **Context:** ~256K tokens
- **Quality:** ★★★★☆ (excellent)
- **Speed:** Fast
- **Best For:**
  - Complex coding tasks
  - Code review
  - Architecture discussions
  - When you need better quality than budget models
- **When to Use:** Good fallback with more budget headroom than K2.6

### Premium Models (Use Sparingly!)

#### GLM-5 — Reasoning Specialist

- **Model ID:** `glm-5`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Cost:** **1,150 requests per $12**
- **Context:** ~200K tokens
- **Quality:** ★★★★☆ (excellent)
- **Best For:**
  - Multi-step reasoning
  - Complex planning
  - Algorithm design
  - Difficult debugging
- **When to Use:** When reasoning/planning is required and budget models fail

#### GLM-5.1 — Maximum Quality

- **Model ID:** `glm-5.1`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Cost:** **880 requests per $12** (most expensive!)
- **Context:** ~200K tokens
- **Quality:** ★★★★★ (best available)
- **Speed:** Moderate
- **Best For:**
  - Critical architectural decisions
  - Complex multi-file refactoring
  - Production code review
  - When you need the absolute best quality
- **When to Use:** Only when cheaper models can't handle the task

### Long Context Specialists

#### MiniMax M3 — Latest Long Context

- **Model ID:** `minimax-m3`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages`)
- **Cost:** **3,200 requests per $12**
- **Context:** **~1M tokens**
- **Quality:** ★★★☆☆
- **Best For:**
  - Very large codebases
  - Long conversations with full context
  - Multi-file analysis
- **When to Use:** Default long-context model (successor to M2.7)

#### MiniMax M2.7 — Proven Long Context

- **Model ID:** `minimax-m2.7`
- **Endpoint:** **Anthropic-compatible** (`/v1/messages`)
- **Cost:** **3,400 requests per $12**
- **Context:** **~1M tokens**
- **Quality:** ★★★☆☆

#### MiMo-V2.5-Pro — Advanced Reasoning

- **Model ID:** `mimo-v2.5-pro`
- **Endpoint:** **OpenAI-compatible** (`/chat/completions`)
- **Cost:** **3,250 requests per $12**
- **Context:** ~128K tokens
- **Quality:** ★★★★☆

## Usage Limits

OpenCode Go limits:

- **5-hour limit:** $12 of usage
- **Weekly limit:** $30 of usage
- **Monthly limit:** $60 of usage

### Cost Comparison Example

**Scenario:** You want to make 5,000 requests this month.

| Model            | Cost | Can you do it?        |
| ---------------- | ---- | --------------------- |
| DeepSeek V4 Flash| ~$2  | ✅ Yes, easily        |
| MiMo-V2.5        | ~$2  | ✅ Yes, easily        |
| MiniMax M2.5     | ~$10 | ✅ Yes                |
| Qwen3.7 Plus     | ~$14 | ✅ Yes                |
| Qwen3.6 Plus     | ~$18 | ✅ Yes                |
| Kimi K2.5        | ~$32 | ❌ Exceeds $30 weekly |
| GLM-5            | ~$52 | ❌ Exceeds limits     |
| GLM-5.1          | ~$68 | ❌ Exceeds limits     |

### Optimizing Your Usage

**Strategy 1: Tiered Approach**

```
1. Start with DeepSeek V4 Flash or Qwen3.7 Plus (cheap, good quality)
2. If it fails, try Kimi K2.5 (better quality)
3. If still failing, use GLM-5 (reasoning)
4. Only for critical tasks: GLM-5.1 or Qwen3.7 Max (premium)
```

**Strategy 2: Task-Based Selection**

```
Background ops (grep, ls, cat) → Qwen3.7 Plus or DeepSeek V4 Flash
General coding → Kimi K2.6 or Qwen3.6 Plus
Complex features → Kimi K2.6
Architecture/Planning → GLM-5
Critical review → GLM-5.1 (rarely)
```

## Fallback Chains for Cost Efficiency

```json
{
  "fallbacks": {
    "background": [
      { "model_id": "qwen3.6-plus" },
      { "model_id": "minimax-m2.5" }
    ],
    "long_context": [
      { "model_id": "minimax-m2.7" },
      { "model_id": "minimax-m2.5" }
    ],
    "default": [
      { "model_id": "mimo-v2.5-pro" },
      { "model_id": "qwen3.6-plus" }
    ],
    "think": [
      { "model_id": "kimi-k2.6" },
      { "model_id": "qwen3.7-plus" }
    ],
    "complex": [
      { "model_id": "glm-5" },
      { "model_id": "kimi-k2.6" }
    ],
    "fast": [
      { "model_id": "qwen3.5-plus" },
      { "model_id": "deepseek-v4-flash" },
      { "model_id": "minimax-m2.5" }
    ]
  }
}
```

**Rule of thumb:** If a task succeeds with a cheap model, it doesn't need an expensive one. Only fall back to expensive models when necessary.

## Quick Reference

| Task Type             | Recommended  | Cost (req/$12) | Fallback       |
| --------------------- | ------------ | -------------- | -------------- |
| Read file, ls, grep   | Qwen3.5 Plus | 10,200         | Qwen3.6 Plus   |
| General coding        | Qwen3.6 Plus | 3,300          | Kimi K2.5      |
| Complex features      | Kimi K2.6    | 1,850          | MiMo-V2.5-Pro  |
| Long context (>80K)   | MiniMax M2.5 | 6,300          | MiniMax M2.7   |
| Reasoning/planning    | GLM-5        | 1,150          | Kimi K2.6      |
| Critical architecture | GLM-5.1      | 880            | GLM-5          |
| Bulk operations       | Qwen3.5 Plus | 10,200         | MiniMax M2.5   |

## Cost-Saving Tips

1. **Use Qwen3.6 Plus as default** — 3,300 req/$12 is plenty for most tasks
2. **Reserve GLM-5.1 for critical tasks only** — 880 req/$12 drains budget fast
3. **Use Qwen3.5 Plus for simple operations** — 10,200 req/$12 is unbeatable
4. **MiniMax M2.5 for long context** — 6,300 req/$12 with 1M context is amazing value
5. **Use Zen free-tier models** for non-critical tasks — deepseek-v4-pro, grok-build-0.1, big-pickle, and others cost $0
6. **Monitor your usage** in the [OpenCode console](https://opencode.ai/auth)

## See Also

- [OpenCode Go Documentation](https://opencode.ai/docs/go/)
- [oc-go-cc Configuration](../configs/config.example.json)
- [README.md](../README.md) for setup instructions

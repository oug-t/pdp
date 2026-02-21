# PDP (Prompt Data Protocol)

The `robots.txt` for AI Prompts.

PDP is a standardized protocol that allows users and developers to define exactly how their prompts and data can be used for AI training, directly via HTTP headers or JSON metadata.

## The Problem

Currently, when a user submits a prompt to an AI service, consent to train on that data is buried in terms of service or global account toggles. There is no standard, granular way to control data usage per request.

## The Solution

PDP introduces a simple 3-tier system passed via the `X-PDP-Level` header:

- `0`: **Strictly Private.** Do not store, do not train.
- `1`: **Local Refinement.** May be used to personalize the current user's session/model only.
- `2`: **Global Training.** Consent granted to train the base model.

## Quickstart (Example Request)

```http
POST /v1/chat/completions HTTP/1.1
Host: api.openai.com
X-PDP-Level: 0
Content-Type: application/json

{
  "model": "gpt-4",
  "messages": [{"role": "user", "content": "Hello!"}]
}
```

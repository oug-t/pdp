# Contributing to PDP

Thank you for your interest in the Prompt Data Protocol. We are currently in **Draft v1** and actively seeking community feedback on the specification structure and header definitions.

## How to Get Involved

### 1. Request for Comments (RFC)

The core specification is located in `spec/pdp-v1.md`.

- **Critique the Protocol:** Open an issue with the tag `[RFC]` if you believe the privacy levels (0-2) are insufficient or if the header name `X-PDP-Level` conflicts with existing standards.
- **Architecture Reviews:** We welcome feedback on edge cases (e.g., streaming responses, WebSocket handling, or conflicting headers).

### 2. Library Implementations

We aim to support the entire AI stack.

- **New Languages:** If you want to port PDP to Rust, Ruby, or Swift, please create a new directory in `libs/` and match the fail-safe behavior of the existing libraries.
- **Improvements:** PRs that reduce dependency bloat or improve type safety in existing libraries are prioritized.

### 3. Pull Request Process

1. Fork the repo and create your branch from `main`.
2. Ensure your code includes the standard fail-safe default (Level 0 / Private).
3. Update the relevant README section if you change public APIs.

## Philosophy

PDP is designed to be **minimalist** and **fail-closed**.

- **Minimalist:** We use standard HTTP headers, not custom JSON payloads.
- **Fail-Closed:** If a header is missing or invalid, the system must default to **Private (Level 0)**.

Join the discussion in [Issues](https://github.com/oug-t/pdp/issues).

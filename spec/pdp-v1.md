# Prompt Data Protocol (PDP) Specification

**Version:** 1.0 (Draft)  
**Status:** Proposed Standard

## 1. Abstract

The Prompt Data Protocol (PDP) defines a standardized method for clients to communicate data usage and training consent to AI service providers. It allows per-request granularity for data retention and model training.

## 2. Transmission Mechanisms

PDP declarations MUST be transmitted via one of two standard methods:

### 2.1 HTTP Header (Preferred)

Clients communicating over HTTP(S) SHOULD include the `X-PDP-Level` header in their requests.

**Syntax:**
`X-PDP-Level: <integer>`

### 2.2 JSON Payload

For WebSocket connections or when modifying headers is impossible, clients MAY include a `pdp_level` key at the root of the JSON payload.

**Syntax:**
`{ "pdp_level": <integer>, "messages": [...] }`

## 3. PDP Levels

Providers implementing PDP MUST respect the following numerical levels:

- **Level `0` (Strictly Private)**
    - **Rule:** The provider MUST NOT store the prompt or response permanently. The provider MUST NOT use the data to train, fine-tune, or evaluate any generalized or user-specific models.
- **Level `1` (Local Refinement)**
    - **Rule:** The provider MAY store the data and use it to personalize the experience or fine-tune models _strictly_ for the current user or organization. The data MUST NOT leak into global base models.
- **Level `2` (Global Training)**
    - **Rule:** The provider MAY use the data for global model training, telemetry, and evaluation, in accordance with their standard terms of service.

## 4. Fallback Behavior

If a request is received without a valid PDP indicator, providers SHOULD default to the consent settings configured at the user's account level. If no account-level setting exists, providers SHOULD default to Level `0` (Strictly Private).

## 5. Extensibility

Future versions of this protocol may introduce a `X-PDP-Options` header for granular flags (e.g., `X-PDP-Options: ttl=3600`), but v1 parsers MUST ignore unrecognized headers to maintain backward compatibility.

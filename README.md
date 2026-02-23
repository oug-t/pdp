# PDP: Prompt Data Protocol
### *The `robots.txt` for the Generative AI Era.*

[**Explore Live Demo →**](https://oug-t.github.io/pdp/)

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE) [![Status](https://img.shields.io/badge/Status-Draft_RFC-orange)](https://github.com/oug-t/pdp/issues)

---

## Overview

Current AI data consent is binary: you accept a provider's global **Terms of Service**, or you don't use the model. There is no middle ground for per-request privacy.

**PDP** is a minimalist, header-based standard for granular, per-prompt data consent. By injecting an `X-PDP-Level` header into HTTP requests, users explicitly signal how their data should be handled—moving consent from a static legal document to a programmable signal.

### Why it matters
* **Granular Control**: Toggle privacy levels on a per-message basis.
* **Fail-Safe**: All libraries default to `Level 0` (Private) if the signal is missing.
* **Non-Invasive**: Lives in the transport layer; zero changes required to prompt JSON payloads.

## The Spec (Draft v1)

| Level | Name | Compliance Requirement |
| :--- | :--- | :--- |
| **0** | **Private** | `NO_STORE`, `NO_TRAIN`. Immediate discard after inference. |
| **1** | **Personal** | `STORE_SESSION`. Allowed for user history and RAG context only. |
| **2** | **Global** | `STORE_PERM`. Consent granted for base model improvements. |

## Quick Start

### Python (AI/ML)

```python
# libs/python/pdp.py
from libs.python.pdp import Level, get_header
import requests

# Send a "Strictly Private" prompt
headers = get_header(Level.PRIVATE)
response = requests.post("[https://api.llm.com/chat](https://api.llm.com/chat)", headers=headers, json={...})
```

### JavaScript / TypeScript (Web)

```javascript
// libs/js/index.js
import { fetchWithPDP } from "./libs/js";

// Standard fetch syntax, but with an extra argument for the PDP Level (0-2).
// Defaults to Level 0 (Private) if omitted.
await fetchWithPDP(
    "/api/chat",
    {
        method: "POST",
        body: JSON.stringify({ prompt: "Hello world" }),
    },
    1
);
```

### Go (Backend Middleware)

```go
// main.go
import "[github.com/oug-t/pdp/libs/go](https://github.com/oug-t/pdp/libs/go)"

func main() {
    // 1. Wrap your router with the middleware
    http.Handle("/api/chat", pdp.Middleware(http.HandlerFunc(handleChat)))
    http.ListenAndServe(":8080", nil)
}

func handleChat(w http.ResponseWriter, r *http.Request) {
    // 2. Extract the safe, typed level from context
    level := pdp.FromContext(r.Context())

    if level == pdp.LevelPrivate {
        // Compliance: Disable logging and DB writes
        log.Println("Processing private request...")
    }
}
```

Java (Enterprise)

```java
// libs/java/PDP.java
import pdp.PDP;

// Immutable map for enterprise http clients
Map<String, String> headers = PDP.getHeader(PDP.LEVEL_GLOBAL);
```


## Philosophy & FAQ

### Is this enforceable?
PDP is a **signal**, not a DRM—much like `robots.txt`. It provides a standardized protocol for compliant providers to honor user intent and enables enterprise proxies to enforce data policies at the network edge.

### Why a header instead of a JSON field?
* **Transport Efficiency**: Gateways can route or block traffic without parsing request bodies.
* **Model Agnosticism**: Keeps payloads "pure" and compatible across providers.
* **Auditability**: Privacy levels can be audited in transit using standard network logs.


## Contributing
We are seeking RFC comments on the specification structure. Please open an issue or submit a PR to `spec/`.

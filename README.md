# PDP (Prompt Data Protocol)

[![Live Demo](https://img.shields.io/badge/Live-Demo_Active-brightgreen)](https://oug-t.github.io/pdp/) [![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)

**The `robots.txt` for AI Prompts.**

PDP is a minimalist, header-based standard that allows users to explicitly signal how their data should be used by AI providers. It moves consent from the "Terms of Service" (global) to the "HTTP Header" (per-request).

## The Spec (Draft v1)

The protocol relies on a single HTTP header: `X-PDP-Level`.

| Level | Name         | Definition                                                                                           |
| :---- | :----------- | :--------------------------------------------------------------------------------------------------- |
| **0** | **Private**  | `NO_STORE`, `NO_TRAIN`. Provider must discard data immediately after inference.                      |
| **1** | **Personal** | `STORE_SESSION`, `TRAIN_USER`. Data can be stored for user history and personal fine-tuning only.    |
| **2** | **Global**   | `STORE_PERM`, `TRAIN_BASE`. Data contributes to the global training set for base model improvements. |

## Language Support

We provide fail-safe libraries for the most common AI stack languages.

| Language       | Path          | Feature                        |
| :------------- | :------------ | :----------------------------- |
| **Go**         | `libs/go`     | Middleware & Context Injection |
| **Python**     | `libs/python` | Client Enums & Header Helpers  |
| **TypeScript** | `libs/ts`     | Type Definitions & Enums       |
| **Java**       | `libs/java`   | Enterprise Static Utilities    |
| **JavaScript** | `libs/js`     | Lightweight Client             |

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

## Contributing

Looking for RFC comments on the specification structure. Please open an issue or submit a PR to `spec/`.

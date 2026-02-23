# Prompt Data Protocol (PDP) Specification v1.0.0

## Abstract

Current AI data consent is binary and account-wide. PDP introduces per-transaction granularity using the `X-PDP-Level` HTTP header.

## Header Definition

Field Name: `X-PDP-Level`
Type: Integer (0-2)

## Compliance Requirements

### Level 0 (Private)

- **Storage:** Volatile memory only.
- **Logging:** No prompts in application logs.
- **Training:** Strictly prohibited.

### Level 1 (Personal)

- **Storage:** Allowed for user history features.
- **Training:** Allowed for LoRA/Fine-tuning specific to this `user_id`.

### Level 2 (Global)

- **Storage:** Permanent.
- **Training:** Allowed for base model pre-training and RLHF.

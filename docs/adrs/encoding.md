---
title: "Encoding"
date:  "2019-01-01"
---

# ADR 02 - Encoding

## Context

The package should support only one human-readable encoding format. But, it may
introduce converters for different representations like Byte representation.

### Scope

#### Base62 vs Hex, Base32, Base36

All other formats except the `Base64` or higher bases, give less space to store
bits into one byte. And since `Base62` only use ASCII alpha-numeric chars to
encode, it still reserves rights for readability.

#### Base62 vs Base64 and above

`Base62` only uses ASCII alpha-numeric chars to represent data which makes it
easy to read, predict the order by a human eye. The rest of the formats above do
not guarantee the simple predictability including `Base64`.

### Benefits

`Base62` encoding provides easy to read, predictable sequences by humans. And
having and supporting only and only one format is increasing compatibility
between systems.

### Technical risks and considerations

Choosing an encoding format that does not support up to `255-bits` with `1 byte`
space is wasting the spaces unnecessarily. Since `Base62` encoding has chosen
for encoding, the package wastes `255 - 62 = 193 bits` per byte(3x space lose).

Case sensitivity; since the `Base62` encoding uses both capital and lower case
ASCII chars, it makes a case sensitivity as a requirement. So any storage system
that stores the `monoton` package sequences MUST provide case-sensitive store.

## Consequences

As known `Base62` encoding limits the usage of spaces inside the integers. At
the same time it gives readability, predictability for human representation of
data. Although, it is possible to support multiple encoders and let users to
choose depending on their needs, the freedom comes with portability and
integration problems between several systems.

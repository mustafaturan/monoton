---
title: "Byte Sizes"
date:  "2019-01-01"
---

# ADR 03 - Byte Sizes

## Context

The total byte size is fixed to 16 bytes for any sequencer. And at least one
byte is reserved to nodes. The package comes with three pre-configured
sequencers and Sequencer interface to allow new sequencers.

### Defaults

The package comes with pre-configured byte sizes for the Nanosecond, Millisecond
and Second sequencers. And it does not allow you to adjust current sizes unless
you create another sequencer. They are adjusted the time and sequence byte sizes
depending on general needs and to increase compatibility between projects.

The current byte sizes:
```
Second:      16 B =>  6 B (seconds)      + 6 B (counter) + 4 B (node)
Millisecond: 16 B =>  8 B (milliseconds) + 4 B (counter) + 4 B (node)
Nanosecond:  16 B => 11 B (nanoseconds)  + 2 B (counter) + 3 B (node)
```

### Flexibility for any new sequencer

It is totally acceptable that you can create new sequencers with other dynamics.
If a microsecond is important for you then you can create Microsecond sequencer.
Also if you need to adjust the maximum available counter or second partition
then you can create a new behavior for Sequencer interface.

## Consequences

Although a strict byte size is limiting the space for nodes and sequences, 16 B
gives enough flexibility for time, counter and nodes. In the next 50 years, it
could be necessary to provide a strategy to upgrade byte size to 32 B.

---
title: "Time"
date:  "2019-01-01"
---

# ADR 01 - Time

## Context

The `monoton` package provides sequences based on the `monotonic` time which
represents the absolute elapsed wall-clock time since some arbitrary, fixed
point in the past. It isn't affected by changes in the system time-of-day clock.

## Consequences

Since the `monotonic` time needs extra calculation steps when it is compared to
regular `system` time, it also consumes an extra time while generating
sequences.

Moreover, according to the documentation of Go language [time package](https://golang.org/pkg/time/),
on some systems, the monotonic clock will stop if the computer goes to sleep. On
such a system, t.Sub(u) may not accurately reflect the actual time that passed between t and u which will result with incorrect sequences.


# Circuit Breaker Pattern in Go

This repository provides a basic Go implementation of the Circuit Breaker pattern, a design pattern used in computer programming to improve the stability and resilience of a system.

## Table of Contents

- [Overview](#overview)
- [Usage](#usage)
- [How it Works](#how-it-works)
- [Example](#example)
  
## Overview

The `CircuitBreaker` struct offers a straightforward approach to the pattern, ensuring that systems don't continuously make calls to a service or operation that's likely to fail.

## Usage

To use the `CircuitBreaker`:

1. Initialize it with a failure threshold and a reset timeout:
   ```go
   cb := NewCircuitBreaker(threshold, resetTimeout)
   ```

2. Make a call through the circuit breaker:
   ```go
   err := cb.Call(yourFunction)
   ```

## How it Works

- **CLOSED**: Initial state. Calls to the service are allowed.
- **OPEN**: If failures exceed the threshold, the circuit transitions to this state. Any call will be denied while in this state, ensuring the system doesn't get overwhelmed.
- **HALF-OPEN**: After the reset timeout passes in the OPEN state, the circuit transitions to this state. Limited requests are allowed to test the system's health.

If the service starts responding successfully again, the circuit breaker will close the circuit and allow regular calls to the service.

## Example

The `main` function demonstrates the Circuit Breaker in action:

1. Makes multiple calls to a simulated failing service.
2. Waits for a while, simulating the reset timeout.
3. Attempts another call to a simulated successful service.
4. Outputs the current state of the circuit breaker.

# Strategy Pattern

## Intent
Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

## Problem Solved
Avoids massive conditional statements (if-else or switch) when multiple variants of an algorithm are needed. Promotes Open/Closed Principle – new strategies can be added without modifying existing context code.

## Participants
- **Context** – holds a reference to a `Strategy` and delegates the algorithm to it.
- **Strategy (interface)** – declares the common operations that all concrete strategies must implement.
- **ConcreteStrategy** – implements the strategy interface with a specific algorithm.

## Structure (UML-like summary)
```
┌─────────────┐         ┌─────────────────┐
│   Context   │────────▶│   «interface»   │
│ - strategy  │         │    Strategy     │
│ + Execute() │         │ + Algorithm()   │
└─────────────┘         └────────┬────────┘
                                 △
                                 │
                ┌────────────────┼────────────────┐
                │                │                │
       ┌────────┴────────┐ ┌──────┴──────┐ ┌───────┴────────┐
       │ ConcreteStrategyA│ │ConcreteStrategyB│ │ConcreteStrategyC│
       │ + Algorithm()    │ │ + Algorithm()   │ │ + Algorithm()   │
       └─────────────────┘ └─────────────────┘ └─────────────────┘
```

## Example in this repository
**Path:** `behavioural/strategy`

The example simulates a payment gateway system where different gateways (Stripe, UPI, PayPal, Credit Card) implement the same `PaymentGateway` interface. The `gateway` context holds a strategy and delegates payment operations to it.

### Key implementation notes:
- **Factory:** `CreatePaymentGateway(gatewayType)` returns `(*gateway, error)` and reports unknown types.
- **Safety:** Public methods return an explicit error when no strategy is configured. `UpdatePaymentGateway` returns an `error` after attempting to replace the strategy.

### Run the example:
```bash
go run ./behavioural/strategy
```

## When to use
- You need to switch algorithm/behavior at runtime (e.g., different pricing strategies, compression algorithms, route calculations).
- You want to isolate algorithms for easier testing and extension.
- You have multiple conditional branches selecting similar behavior.

## Benefits
- Improves code organization, testability and adherence to Open/Closed Principle.
- New strategies can be added without changing the context code.
- Encourages composition over inheritance.

## Drawbacks & considerations
- Slight increase in number of objects/classes.
- Clients must be aware of strategy creation; prefer a registry or factory to centralize creation.

## Go-Specific Suggestions
- Export the `PaymentGateway` interface if external packages should implement new gateways.
- Consider a registry map `map[GatewayType]func() PaymentGateway` to avoid modifying factory switch statements.
- Add unit tests for factory error cases, `UpdatePaymentGateway`, and each concrete gateway.
- Use functional options or constructor injection to supply the initial strategy.

## Real-World Analogy
A navigation app: the route calculation algorithm (car, walking, public transport) is the strategy. The context (the navigator) delegates route calculation to the chosen strategy without knowing its details.
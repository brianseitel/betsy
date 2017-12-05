# betsy

A feature flag library

![betsy ross](https://www.wikitree.com/photo.php/thumb/e/e9/Betsy-Ross.jpg/300px-Betsy-Ross.jpg)

## Basic Example

Instantiate `FeatureFlag`

```go
flags := NewFeatureFlags()
```

Create a Rule

```go
r, _ := NewRule(true)
flags.Add("bool", r)
```

Check whether rule is enabled for a value

```go
flags.Allowed("bool", true) // returns true
flags.Allowed("bool", false) // returns false
```

Other rule types:

```go

// An integer rule
r, _ = NewRule(1)
flags.Add("int", r)

flags.Allowed("int", 1) // passes (true)
flags.Allowed("int", 2) // fails (false)

// A string rule
r, _ = NewRule("ice cream")
flags.Add("string", r)

flags.Allowed("string", "ice cream") // passes (true)
flags.Allowed("string", "football") // passes (false)

// A slice of integers rule
r, _ = NewRule([]int{1, 2, 3, 4, 5})
flags.Add("metros", r)

// See if value is in slice of integers
flags.Allowed("metros", 1) // passes (true)
flags.Allowed("metros", 15) // fails (false)

// A slice of strings rule
r, _ = NewRule([]string{"banana", "apple", "orange"})
flags.Add("fruit", r)

// See if values are in slice of strings
flags.Allowed("fruit", "banana") // passes (true)
flags.Allowed("fruit", "durian") // fails (false)
```

Display all flags and their values

```go
flags.List()
```

## Rules

Rules are created by using the `NewRule(input interface{}) (Rule, error)` method. Supported types are:

* int
* string
* bool
* []int
* []string

If an invalid type is provided, it will return a `NoopRule` and an error. A `NoopRule` will always fail.

## Allowed / Denied

The primary flag method is `Allowed(name string, value interface{}) bool`. Depending on the rule type, this method compares the `value` with the acceptable values in the `Rule`. If the value matches, then it passes. If not, it fails.

There is also a `Denied` method for syntactic sugar, but it's just an alias to `!Allowed()`.

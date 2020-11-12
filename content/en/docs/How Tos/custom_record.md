---
title: "Create a custom Record"
date: 2020-10-29
weight: 300
description: >
  Record is the interface provided by Baker to represent an "object" of data
---

As you can read in the [Record and LogLine page](/docs/core-concepts/record_implementation/),
Baker processes objects in form of “records”.

A `Record`, in Baker, is an interface that provides an abstraction over a record of flattened data,
where columns of fields are indexed through integers.

If the Record implementations provided by Baker doesn't fit your needs, you can create your own
version of the Record, implementing the
[`Record` inteface](https://pkg.go.dev/github.com/AdRoll/baker#Record).

## How to use a custom version of the Record

Once your Record version is ready, you need to use it in your code.

In order to do so, some functions may be implemented while instantiating
[`baker.Components`](/docs/how-to/baker_components/):

```go
type Components struct {
	Validate      ValidationFunc
    CreateRecord  func() Record
    FieldByName func(string) (FieldIndex, bool)
    FieldName   func(FieldIndex) string
    //... other functions
}
```

### Validate

`Validate` is the function used to validate a record. It is called for each processed record
unless not set or when the `[general] dont_validate_fields = true` configuration is set in
the TOML file.

Regardless of the TOML configuration, the function is passed to all components that can use
it at their will.

### CreateRecord

`CreateRecord` is the function that creates a new record. If not set, a default function is
used that creates a `LogLine` with `,` as field separator.

The function is used internally by Baker to create new records every time a new one comes from
the input.

The function is also passed to components that can use it to create new records while processing.

### FieldByName

`FieldByName` gets a field index by its name. The function is mainly used by the components
(that receive it during setup) to retrieve the index of a field they need for filtering or
processing, but it is also used internally by Baker when sending fields to the output
(when at least one field is selected in the output TOML configuration).

### FieldIndex

`FieldName` gets a field name by its index. The function is passed to components that can use
it for their internal logic.

## RecordConformanceTest

The `test_helper.go` provides a `RecordConformanceTest` test helper whose goal is to give the
user a structured test for new implementations of the Record.

The helper receives the implementation of `CreateRecord` and creates new records testing
them against a set of requirements.

{{% alert title="Warning" color="warning" %}}
The conformance test provides a way to verify that a record implementation respects the
invariant that Baker requires for a Record implementation and thus it should always
be executed against all custom implementations of the Record.
{{% /alert %}}
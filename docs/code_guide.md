# Code Guide

This document serves as a guide for how code is written in the freecreate-go codebase.

It covers project structure as well as code philosophy, style, and best practices.

## Errors

Due to golangs lack of stack tracing with its errors, FreeCreate uses a custom error type that has logging enabled. It can be viewed directly in the `err` package within the `internal` directory.

### Returning Errors

All functions that return errors should return this custom error.

```go
func myFunc() err.Error{
    return err.Error{}
}
```

### Handling Native Errors

Whenever a native error is generated, it should be converted to a custom error using the `err.NewFromErr` function, which takes in an `error` as its argument:

```go
sErr := //some code that generates a native error
if sErr != nil {
    e := err.NewFromErr(sErr)
    // log or return the custom error
}
```

### Treat Errors as Values

code should follow the same "error as value" paradigm that is native to go. Instead of checking if a native error is `nil`, check if the custom error's `E` property is `nil`:

```go
if cErr.E != nil {
    // handle the error
}
```

### Error Naming

when creating error variables, never call them `err`, as that conflicts with the package name. call them `e`, or prepend a letter or two to `err` based on the operation.

For example, let's say a `sum` function might return an error:

```go
val, sErr := sum()
```

we have named our error value `sErr` because the function generating it is called `err`.

This is a good practice to follow regardless, as it provides a standard convetion for functions where multiple errors might be generated.

### Logging Errors

If an error reaches the top of its stack - i.e. it is not being returned by the function in which it appears - call the `Log` function to log its stack trace.

## MQC Architecture - Model, Query, Controller

While the `view` portion of the application is handled by a separate frontend codebase, code is still split up along 3 paradigms - Models, Queries, and Controllers.

Code has been separated along these three paradigms to make operations more modular, standardized, and testable.

### Models

Models are responsible for describing the "shape" and type of data and validating data. They represent individual data entities that exist within the database.

The primary model must have all fields and only those fields that correspond to data stored by an individual entity of that type within the database. Ex: a `User` model must have all fields and only those fields that an individual database record will store about a User.

Models are implemented as `structs`.

Any data that pertains to a specific type of data will be implemented as a model-related struct.
This includes data the API receives and emits.

All structs pertaining to a specific type of data should be stored in the model file.

All models that will be entered into the database must have a factory function to generate the model and a validation function that validates the model's data upon generation.

All generator and validation functions must have tests.

### Queries

Queries are database queries that are run for given operations.

Each individual operation is given its own query file, whose name is the name of the query. If multiple queries are run within the same operation, they will be included in the same query file.

Factory functions are used to generated database query strings and / or filters as well as requisite query params.

A query file also contains a function that runs the actual query.

All factory functions for query strings and filters as well as query params must have their own tests checking their validity.

Neo4j query strings must be constructed using validated labels. See [Neo4j labels](#neo4j-labels) for more detail.

Tests that test query string validity should <em>not</em> use validated labels - instead write the desired query by hand.

### Controllers

Controllers (called `handlers` in the codebase) are responsible for handling API endpoints. They are essentially the "glue" that handles incoming requests, call upon models and queries to format and validate data and run queries, and then send back a corresponding response.

## Tests

### Unit tests

All code files should have a corresponding unit test file. Any function that does not require connection to outside resources should be unit tested.

Unit test files are named identically to their corresponding code file, and are stored in the same package immediately adjacent to the code file. This is to ensure coverage of functionality internal to the package without exporting code that is only used internally.

## Neo4j Labels

Due to Neo4j's flexible schema, we cannot limit the types of nodes that are entered into the database using Neo4j itself (at time of this writing).

For those familiar with SQL databases, this is a paradigm shift.

Labels are also case sensitive - a `User` and `USER` label are registered differently by neo4j.

To ensure that typos do not accidentally inject invalid labels into the database, always use labels generated from label factory functions.



// Technical Features of Go
// First-class concurrency primitives
// Type safety enforced by compiler
// Memory safe; no use-after bugs; Garbage collection
// Compiles to machine code

// Packages & Modules
// Programs are written as one or more packages
// Packages should be focused and perform a single thing
// - Argument parsing
// - Drawing Graphics
// - Handling HTTP requests

// Modules are a collection of packages
// Created by having a go.mod file in the root directory of your project
// All Go projects should have a go.mod file

// all primitive data types are numeric ie. a stream of bytes

// Type Aliases
// Useful for providing indication of what kind of data is being utilized
// type UserId int
// type Speed float64
// type Velocity Speed
// Converting between types can be done with parentheses

// Strings in Go are actually an array of bytes and a string length
// A string is the data type for storing multiple runes
// When iterating a string, iteration occurs over bytes
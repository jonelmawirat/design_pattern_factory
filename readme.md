# Factory Design Pattern in Go

This repository contains my implementation of the **Factory Design Pattern** in Go (Golang). The code is based on the concepts and examples presented in the book [Go Design Patterns](https://learning.oreilly.com/library/view/go-design-patterns/9781786466204/ch02s03.html).

## Overview

The **Factory Design Pattern** is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. This pattern helps encapsulate the instantiation logic and promotes loose coupling.

### Key Points:
- Simplifies object creation by centralizing the logic.
- Helps in selecting the appropriate subclass based on input or configuration.
- Promotes abstraction and makes the code more modular.


### `factory.go`
This file contains the factory function. The function returns different implementations based on the provided input.

### `products.go`
Here, I defined various concrete types that implement the same interface. These represent the "products" created by the factory function.

## Prerequisites

- Go 1.23+ installed on your machine
- Basic knowledge of Go and design patterns

## Usage

To run the code, use the following command:

```bash
go test -v ./...
```

This will execute the code and show how the factory function dynamically creates objects based on the input.

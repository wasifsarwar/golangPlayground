# Golang Notes

In golang:
- every value has a type
- every function has to specify the type of its arguments

### Interface
In Go, interfaces are a powerful way to define behavior. They allow us to specify a set of methods that a type must implement, without dictating how those methods should be implemented. This promotes flexibility and decouples the code, enabling polymorphism.

Key Concepts of Interfaces in Go:
- Definition: An interface is defined using the type keyword followed by the interface name and the method signatures it requires. For example:
```
   type Greeter interface {
       getGreeting() string
   }
```

- Implementation: Any type that implements the methods defined in an interface satisfies that interface. For instance, both `englishBot` and `spanishBot` implement the `getGreeting` method:

```
   func (eb englishBot) getGreeting() string {
       return "Hello There"
   }

   func (sb spanishBot) getGreeting() string {
       return "Hola"
   }
```

- Using Interfaces: You can use interfaces to write functions that can accept any type that implements the interface. For example, if you had a function that takes a `Greeter` interface:
```
   func printGreeting(g Greeter) {
       fmt.Println(g.getGreeting())
   }
```

- Polymorphism: This allows you to write more generic code. You can pass any type that implements the Greeter interface to the printGreeting function:
```
   eb := englishBot{}
   printGreeting(eb) // Outputs: Hello There

   sb := spanishBot{}
   printGreeting(sb) // Outputs: Hola
```   

In golang, there are concrete types and interface types
- Concrete types: 
    - built in types (`int`, `float`, `structs`, `map`)
    - Specified, defined type that has a fixed structure and behavior
    - memory is allocated for an instance of a type
    - can create variables of concrete types and use them directly


- Interface types: 
    - An interface type defines a set of method signature that a concrete type MUST implement
    - It doesn't provide any implementation
    - It specifies behavior without dictating how that behavior is implemented
    - Allows polymorphism by allowing functions to accept any type that implements the interface


Rules of interfaces:

## **Variable Declaration**
In Go, we **declare a variable** of a certain **type**:
- A variable can only hold **values** of its declared **type**.
- Go is a **statically typed** language.

---

## **Basic Types / Built-in Types / Primitive Types**
  - Data types provided by a programming language as a basic building block. Most languages allow more complicated composite types to be constructed starting from these basic types.
  - Data types for which the programming language provides built-in support.
  - In most programming languages, all basic data types are built-in.
  - For more details, refer to the [Go Built-in Package Documentation](https://pkg.go.dev/builtin).

---

## **Aggregate Types**
Aggregate types combine multiple values together.

### Examples:
- **Array**: A fixed-size collection of elements of the same type.
- **Slice**: A dynamically sized, flexible view into an array.
- **Struct**: A collection of fields that can hold values of different types.

### Composite / Compound / Struct Types:
- Include **values of different types**, e.g., `struct`.
- **Composition** is the act of constructing a `struct` (a composite type) by:
  - Combining many different **types** into one structure.
  - Creating a **data structure** that holds values of multiple **types**.
- **Combining smaller types** into a larger type.
- **Embedding types** and **inner-type promotion**:
  - Enables a struct to **inherit fields and methods** of an embedded type.

---

## **Go Specification Types**
Go provides a variety of built-in types, as per the [Go Specification](https://go.dev/ref/spec):
1. **Boolean**
2. **Numeric**  
3. **String**
4. **Array**
5. **Slice**
6. **Struct**
7. **Pointer**
8. **Function**
9. **Interface**
10. **Map**
11. **Channel**

---

## **Conclusion**
Go's robust type system ensures clarity, consistency, and safety in variable usage and data structuring. Whether you’re using basic types or aggregating multiple types, Go's static typing and comprehensive specifications make it a powerful language for efficient programming.
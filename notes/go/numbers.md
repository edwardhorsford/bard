# Numbers

In Go, there are severl different number types, though they are usually split
into one of two categories: integers and floating-point numbers.

## Integers

Any number without a decimal point.

Go has several integer types to represet the different bit sizes that the 
number should use:

| Type   |                                    |
|--------|------------------------------------|
| uint8  | Unsigned Integer 8 bits (1 byte)   |
| uint16 | Unsigned Integer 16 bits (2 bytes) |
| uint32 | Unsigned Integer 32 bits (4 bytes) |
| uint64 | Unsigned Integer 64 bits (8 bytes) |
| int8   | Signed Integer 8 bits (1 byte)     |
| int16  | Signed Integer 16 bits (2 bytes)   |
| int32  | Signed Integer 32 bits (4 bytes)   |
| int64  | Signed Integer 64 bits (8 bytes)   |


Unsigned integers only contain positive numbers (or zero). In addition there two 
alias types: byte which is the same as uint8 and rune which is the same as 
int32. Bytes are an extremely common unit of measurement used on computers 
(1 byte = 8 bits, 1024 bytes = 1 kilobyte, 1024 kilobytes = 1 megabyte, etc.) 
and therefore Go's byte data type is often used in the definition of other 
types. There are also 3 machine dependent integer types: uint, int and uintptr. 
They are machine dependent because their size depends on the type of 
architecture you are using.

Generally if you are working with integers you should just use the int type.

## Floating-point numbers

Any number with a decimal.

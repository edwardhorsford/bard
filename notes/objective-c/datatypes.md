# Objective-C Data Types

## NSString

```objective-c
NSString *myString = @"meerkat manor"
```
* immutable
* support for Unicode and UTF-8 characters

### stringWithFormat

```objective-c
int age = 15;
NSString *formattedString = [NSString stringWithFormat:@"I am @d years old", age]
```

# Arrays

```ruby
new_array = [1, 2, "kittens", true, [1,2,3], { kittenSound: 'meow' }]
```

An array type is a data type that is meant to describe a collection of
elements (values or variables), each selected by one or more indices
(identifying keys) that can be computed at run time by the program. It can store
any data type, such as strings, numbers, hashes, and other arrays.

The values of an array are accessed by their index number, like so:

```ruby
irb(main)> new_array[2] // 'kittens'
```

**Note: Array indexes start at 0**

You can also access a range of values in an array using a range, like so:

```ruby
new_array[0..2] // [1,2,'kittens']
```

## Array Methods

### push or <<
- adds items to an array


```ruby
irb(main)> array_name = [1,2,3]
irb(main)> array_name.push(4)
=> [1,2,3,4]
irb(main)> array_name << 5
=> [1,2,3,4,5]
```

----

### empty?

```ruby
irb(main)> array_name = []
irb(main)> array_name.empty?
=> true
irb(main)> array_name = [1,2,3]
irb(main)> array_name.empty?
=> false
```

This method will return true if the array is empty and false if the array has
values.

-----

### pop

```ruby
irb(main)> array_name = [1,2,3]
irb(main)> array_name.pop
=> [1, 2]
```

This method removes the last item from an array and returns the removed value.

----

### shift

```ruby
irb(main)> array_name = [1,2,3]
irb(main)> array_name.shift
=> [2,3]
```

This method removes the first item from an array and returns the removed value.

----

### unshift

```ruby
irb(main)> array_name = [2,3]
irb(main)> array_name.unshift(1)
=> [1,2,3]
```

This method adds items to the beginning of an array.
----

### sort

```ruby
irb(main)> array_name = [3,2,1]
irb(main)> array_name.sort
=> [1,2,3]
irb(main)> array_name
=> [3,2,1]
irb(main)> array_name.sort!
=> [1,2,3]
irb(main)> array_name
=> [1,2,3]
```

This method returns an array with the values of the original array sorted from
lowest to greatest. If you want to make a permanent change to the original array,
add an '!' to the end.

----

### shuffle

```ruby
irb(main)> array_name = [1,2,3,4,5]
irb(main)> array_name.shuffle
=> [3,2,4,1,5]
```

This method returns an array with the values of the original array moved around in a random order

----

### reverse

```ruby
irb(main)> array_name = [1,2,3]
irb(main)> array_name.reverse // [3,2,1]
irb(main)>array_name // [1,2,3]
irb(main)> array_name.reverse!
irb(main)> array_name // [3,2,1]
```

This method returns an array with the values of the original array in reverse
order. If you want to make a permanent change to the original array, add an '!'
to sort.

----
### join

```ruby
irb(main)> array_name = ['the','cat','is','awesome']
irb(main)> array_name.join(' ') // 'the cat is awesome'
```

This method joins the values of the array into a string delimited by the value
passed to this method.

----

### any?

```ruby
irb(main)> array_name = [false, true, false]
irb(main)> array_name.any? // true
```

This method returns true if any of the values within the array are truthy.

----

### all?

```ruby
irb(main)> array_name = [false, true, false]
irb(main)> array_name.all? // false
irb(main)> array_name = [true, true, true]
irb(main)> array_name.all? // true
```

This method returns true if all of the values in the array are truthy

----

### each

```ruby
irb(main)> new_arr = []
irb(main)> arr = [1,2,3]
irb(main)> arr.each do |x|
irb(main)> 	new_arr << x += 10
irb(main)> end
=> [1, 2, 3]
irb(main)> new_arr
=> [11, 21, 31]
```

This method iterates over each item in the array executing a block of code on
each item. It also returns the original array.

----

### map

```ruby
irb(main)> arr = [1,2,3]

irb(main)> arr.map! do |x|
irb(main)>   x += 10
irb(main)> end
irb(main)> arr
=> [11, 21, 31]
```

This method iterates over each item in the array executing a block of code on
each item. Unlike the each method, this method returns an array with the values
updated, if the values were modified within the block of code passed to it. You
can make permanent changes to the original array by using a '!'.

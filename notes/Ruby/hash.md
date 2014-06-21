# Hash

```ruby
hash_name = { president: 'Barack Obama' }
```

A hash table (also hash map) is a data structure used to implement an
associative array, a structure that can map keys to values. A hash table uses
a hash function to compute an index into an array of buckets or slots, from
which the correct value can be found.

## Accessing Values

```ruby
hash_name[:president] // 'Barack Obama'
```

The values inside of a hash are accessed using the keys as Ruby symbols.

## Adding values

```ruby
hash_name['first_lady'] = 'Michelle Obama'
hash_name // { president: 'Barack Obama', first_lady: 'Michelle Obama' }
```

You can add a key-value pair to hash in a similar way as you would declare a
variable, but you must define what the key will be.

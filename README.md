# nx
> Next core for ruby version.

## install:
```rb
# add to Gemfile
gem 'nx', git: 'https://github.com/afeiship/nx'
```

## usage:
```rb
## add bundle/setup:
require "bundler/setup"
require "nx"

h = { "b" => { "c" => [0, [1, 3]] } }
p h.get("b.c[1][1]")
```

## resouces:
- https://bugs.ruby-lang.org/issues/13179
- http://ruby-doc.org/core-2.3.0_preview1/Hash.html#method-i-dig
- https://tagmycode.com/snippet/8218/get-set-hash-values-using-key-path-separated-by-dots#.XIH69BMzaM4
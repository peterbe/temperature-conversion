# temperature-conversion

This is the code for this blog post series:

1. [Converting Celsius to Fahrenheit with Python](https://www.peterbe.com/plog/converting-celsius-to-fahrenheit-with-python)
1. [Converting Celsius to Fahrenheit with TypeScript](https://www.peterbe.com/plog/converting-celsius-to-fahrenheit-with-typescript)
1. [Converting Celsius to Fahrenheit with Go](https://www.peterbe.com/plog/converting-celsius-to-fahrenheit-with-go)
1. [Converting Celsius to Fahrenheit with Ruby](https://www.peterbe.com/plog/converting-celsius-to-fahrenheit-with-ruby)
1. [Converting Celsius to Fahrenheit with Crystal](https://www.peterbe.com/plog/converting-celsius-to-fahrenheit-with-crystal)
1. [Converting Celsius to Fahrenheit with Rust](https://www.peterbe.com/plog/converting-celsius-to-fahrenheit-with-rust)

## Python

```sh
python conversion.py
```

## Ruby

```sh
ruby conversion.rb
```

## TypeScript (Bun)

```sh
bun run conversion.ts
```

## JavaScript (Node)

```sh
node conversion.js
```

## Go

```sh
# go run conversion.go  # debugging
go build -o conversion-go conversion.go
./conversion-go
```

## Crystal

```sh
# crystal conversion.cr  # debugging
crystal build -o conversion-cr conversion.cr
./conversion-cr
```

## Rust

```sh
rustc build -o conversion-rs conversion.rs
./conversion-rs
```

## Deno

```sh
deno run conversion.ts
```

## Benchmark

```sh
hyperfine --shell=none --warmup 3 "python3.12 conversion.py" "bun run conversion.ts" "ruby conversion.rb" "node conversion.js" "./conversion-go" "./conversion-cr" "./conversion-rs" "deno run conversion.ts"
```

Results of `hyperfine` benchmark:

```text
Summary
  ./conversion-rs ran
    1.31 ± 1.30 times faster than ./conversion-go
    1.88 ± 1.33 times faster than ./conversion-cr
    7.15 ± 4.64 times faster than bun run conversion.ts
   14.27 ± 9.48 times faster than python3.12 conversion.py
   18.10 ± 12.35 times faster than node conversion.js
   67.75 ± 43.80 times faster than ruby conversion.rb
```

But note, `hyperfine` will warn with

>  Warning: Statistical outliers were detected. Consider re-running this benchmark on a quiet system without any interferences from other programs. It might help to use the '--warmup' or '--prepare' options.

This is because each program starts and finishes "too fast".
So the benchmark merely becomes a test of **how fast the executable can start**.
Granted, it's a useful thing to know, too.

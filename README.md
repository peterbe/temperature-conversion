# temperature-conversion

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

## Benchmark

```sh
hyperfine --warmup 3 "python3.12 conversion.py" "bun run conversion.ts" "ruby conversion.rb" "node conversion.js" "./conversion-go" "./conversion-cr" "./conversion-rs"
```

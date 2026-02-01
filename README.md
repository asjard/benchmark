- The code in this respositry was adapted from [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark)
- Dependencies for each framework have been isolated and decoupled.
- Test results obtained under the following environment

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz, 6 cores
memory: 32G
```

## Build all frameworks

```bash
bash test-build.sh
```

## Running the Benchmarks

```sh
rm -rf *.png && bash test-latency.sh
```

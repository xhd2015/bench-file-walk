# Benchmark filepath.WalkDir
The original purepose is to compare performance of `filepath.WalkDir` with a custom concurrent implementation using channel, however it seems that `filepath.WalkDir` out performs.

Benchmark data:

|Max Depth|Total Dirs|Cost|Avg Cost/Dir|
|-|-|-|-|
|20|100|4.73ms|0.0473ms|
|20|250|12.90ms|0.0516|
|20|2500|123.10ms|0.04924ms|
|20|5000|251.53ms|0.0503ms|
|20|10000|544.53ms|0.05445ms|

The concurrent implementation is not provided yet.

# How to reproduce?

## Step 1: Prepare the test dirs
```sh
for i in 100 250 2500 5000 10000; 
    do go run ./ genTree tmp_20_$i 20 $i;
done
```
## Step 2: Run the benchmark
```sh
go test -bench=. -benchtime=10s -run=NONE -v
```
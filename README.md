# pixel-count-perf

A random collection of programs that implement a function that given a file with a raw RGBA8 buffer of data, returns the number of pixels that have all bits set (R == 0xFF, G == 0xFF, B == 0xFF and A == 0xFF).

This should perhaps be elaborated to include finding pixels set with an arbitrary color.

Each folder has a different implementation.  Currently the `go-basic` implementation acts as a "baseline" of sorts for both correctness as well as performance (ns of execution time).

# TODO:

1. Better rig to run all flavors of experiments vs filtering when developing one method.
2. CUDA stuff needed to implement first shot algo (how do we measure other things like CPU and memory util?).

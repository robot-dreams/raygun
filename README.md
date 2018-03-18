# raygun

I implemented a very basic ray tracer by following along with the book [Ray Tracing in One Weekend](http://in1weekend.blogspot.com/2016/01/ray-tracing-in-one-weekend.html) this past weekend.  [This reference](https://www.scratchapixel.com/lessons/3d-basic-rendering/introduction-to-shading/reflection-refraction-fresnel) was also extremely helpful for understanding the details of refraction.

I used all of the ideas from the book, but for the most part I tried to avoid looking at the book's code until after I came up with my own approach for a given idea.  Furthermore, although the book uses C++, I used go for my implementation (just because of preference / familarity).

## Usage

If you don't have go installed yet, you can [download the latest version of go from the official website](https://golang.org/dl/).  On Mac OS X, you will also need a recent version of Xcode and the command line tools.

Note that the output format is `ppm` (a simple text format for pixel data), NOT `png`!

```
# Fetch dependencies.
go get github.com/go-gl/mathgl/mgl32

# Clone repo and build everything.
git clone https://github.com/robot-dreams/raygun.git; cd raygun; go build

# Do ray tracing!
./raygun > output.ppm; open output.ppm
```

## Sample Output

The default output parameters (set in `main.go`) are small and meant for rapid iteration (i.e. a few seconds per run).  Using the following larger parameters produces a nicer looking image (but took about 35 minutes to run on a MacBook Pro):

- `nx := 1600`
- `ny := 800`
- `subpixelSamples := 1000`

![Sample Output](https://github.com/robot-dreams/raygun/raw/master/sample.png)

## Future Improvements

- Cleanup
    - Move `main` to separate package
    - Save to other image formats (e.g. png) directly
    - Load a scene from file
    - Actually add some tests, maybe?
- Functionality
    - More types of scene objects (e.g. triangles)
    - Nested indexes of refraction
    - Light sources
    - Textures
- Performance
    - Profile / optimize
    - CPU parallelism
    - Look into using GPU acceleration

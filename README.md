# raygun

I implemented a very basic ray tracer by following along with the book "Ray Tracing in One Weekend" this past weekend.  I used all of the ideas from the book, but for the most part I tried to avoid looking at the book's code until after I came up with my own approach for a given idea.  Furthermore, although the book uses C++, I used go for my implementation (because of preference / familarity).

## Usage

```
go get github.com/go-gl/mathgl/mgl32
git clone https://github.com/robot-dreams/raygun.git
cd raygun
go build
./raygun > output.ppg
open output.ppg
```

## Sample Output

Parameters for sample output (~35 min runtime on a MacBook Pro)

- `nx := 1600`
- `ny := 800`
- `subpixelSamples := 1000`

![Sample Output](https://github.com/robot-dreams/raygun/raw/master/sample.png)

## Future Improvements

- Cleanup
    - Move `main` to separate package
    - Save to other image formats (e.g. png) directly
    - Load a scene from file
- Functionality
    - More types of scene objects (e.g. triangles)
    - Nested indexes of refraction
    - Light sources
    - Textures
- Performance
    - Profile / optimize
    - CPU parallelism
    - GPU acceleration

# Directory Structure
This project should follow [golang's recommended project layout](https://github.com/golang-standards/project-layout) (or at least close to it)

## The gist of it
```
./assets - anything that we need to include (e.g. spritesheets and audio files)
./cmd - anything executable
./docs - documentation
./internal - where most internal packages exist
./pkg - anything that could be used externally by some other project
./.dist - where we build to (should not be included in source control)
./.github - automated test runners
```
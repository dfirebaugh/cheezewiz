# flatbuffers high level
[official docs](https://google.github.io/flatbuffers/)

It's a way to serialize/deserialize data.
Kind of like json.

The big difference is that it get's pressed down into a nice byte array.  This makes transport a lot quicker.
This requires you to write a schema ahead of time.  

## write a schema
checkout the `internal/models/schemas` directory for an example.


## automatically generate some code
you will need the flatc compiler installed

on ubuntu:
```bash
apt install flatbuffers-compiler
```

### example

```bash
flatc --go --grpc -o ./internal/models/ ./internal/models/schemas/greet.fbs
```

> note: you only need to do this if you make a change to the schema.  Once the code is generated, you won't need to change it.
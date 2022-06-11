# ecs adapter
The intention of of the adapter is to provide an interface that will allow us to hotswap the entity component library.

The `adapter.go` file contains an interface.  other files in the package should implement that interface. Ideally the app wouldn't know which ecs solution we are using.


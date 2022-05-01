# pubsub
currently we are using a minimal pubsub packge that lives in the pkg dir.
We may upgrade to a 3rd party solution if we need to.

The most notable deficiency with this solution is that we can't cancel a message once it's sent.  We could potentially look at using something from go's [contetx](https://pkg.go.dev/context) package to implement some form of canceling.
We could also make improvements on idempotency.

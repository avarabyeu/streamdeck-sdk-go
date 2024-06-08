# Stream Deck SDK for Golang

An implementation of Elgato's [Stream Deck SDK](https://developer.elgato.com/documentation/stream-deck/sdk/overview/) to enable writing simple, lightweight plugins in Go.

Credits and Licensing
--------

Events are handled in a similar way to [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo), and in fact uses the same base code. This is the most reasonable way to handle events in this scope.

Usage and importing
-----------

This package requires `github.com/valyala/fastjson`. Sometimes, this will be installed for you or resolved by go module/import managers.

```
go get github.com/OmegaRogue/streamdeck-sdk-go
```

Examples
-------

See the examples directory for a basic example
# BuddhoIO Golang Ext

This package contains dependency free language extensions.

The lang package contains building blocks such as Either, Option and Tuple. These patterns help to reduce code duplication
and improve function reusability. In general these patterns are used to handle errors, optional values and multiple return values
and are common in functional programming languages.

The slices package adds additional functionality to the built-in slice type. This includes but not limited to generic map, filter, flatmap and flatten.

The iter package contains extensions to the golang iter.Seq type. This includes but not limited to generic map, filter, flatmap and flatten.
Additionally, it contains adapters to convert things like channels, grpc streams and sql rows to iter.Seq. These extensions
are especially helpful when composing multiple data sources into a single iter.Seq.

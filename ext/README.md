<!--
  ~ Copyright 2025 BuddhoIO
  ~
  ~ Licensed under the Apache License, Version 2.0 (the "License");
  ~ you may not use this file except in compliance with the License.
  ~ You may obtain a copy of the License at
  ~
  ~     http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing, software
  ~ distributed under the License is distributed on an "AS IS" BASIS,
  ~ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  ~ See the License for the specific language governing permissions and
  ~ limitations under the License.
-->

# BuddhoIO Golang Ext

This package contains dependency free language extensions.

The lang package contains building blocks such as Either, Option and Tuple. These patterns help to reduce code duplication
and improve function reusability. In general these patterns are used to handle errors, optional values and multiple return values
and are common in functional programming languages.

The slices package adds additional functionality to the built-in slice type. This includes but not limited to generic map, filter, flatmap and flatten.

The iter package contains extensions to the golang iter.Seq type. This includes but not limited to generic map, filter, flatmap and flatten.
Additionally, it contains adapters to convert things like channels, grpc streams and sql rows to iter.Seq. These extensions
are especially helpful when composing multiple data sources into a single iter.Seq.

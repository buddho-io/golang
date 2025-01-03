# Copyright 2025 BuddhoIO
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

default: test

# Project name
PROJECT := ext

PUBLISH_LIBRARY := 1

# Go package for this project
PKG := $(shell sed -ne 's/^module //gp' go.mod)

include ../scripts/common-general.mk
include ../scripts/common-go-test.mk
include ../scripts/common-go-tools.mk

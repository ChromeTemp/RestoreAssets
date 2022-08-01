<!-- cspell:word bindata -->

# RestoreAssets [![CI](https://github.com/ChromeTemp/RestoreAssets/actions/workflows/CI.yml/badge.svg)](https://github.com/ChromeTemp/RestoreAssets/actions/workflows/CI.yml)

> A simple library to restore Go assets from "embed" module

## Why?

Coming from go-bindata, it's not so easy to restore assets from "embed" module, since doesn't provide the go-bindata function `RestoreAssets`.

This library tries to solve this problem.

## Usage

```go
package main

import (
    "github.com/ChromeTemp/RestoreAssets"
    "embed"
)

//go:embed assets/*
var assets embed.FS

func main() {
    // Restore all assets to "app_data" directory
    RestoreAssets.From(&assets, "app_data")
}
```

> **Note:** this only works for folders embedded with "embed" module and not for files.

---

<p align="center">Copyright (c) 2022 <a href="https://github.com/ChromeTemp">ChromeTemp</a>, released under the <a href="https://github.com/ChromeTemp/ChromeTemp/blob/main/LICENSE">MIT License</a></p>
<p align="center">ChromeTemp by <a href="https://github.com/Bellisario">Giorgio Bellisario</a></p>

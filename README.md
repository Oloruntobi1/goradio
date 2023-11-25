[![Go Reference](https://pkg.go.dev/badge/github.com/neptunsk1y/goradio.svg)](https://pkg.go.dev/github.com/neptunsk1y/goradio)
[![Test](https://github.com/neptunsk1y/goradio/actions/workflows/test.yml/badge.svg)](https://github.com/neptunsk1y/radiorecord/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/neptunsk1y/goradio)](https://goreportcard.com/report/github.com/neptunsk1y/goradio)
[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Get%20over%20170%20free%20design%20blocks%20based%20on%20Bootstrap%204&url=https://www.froala.com/design-blocks&via=froala&hashtags=bootstrap,design,templates,blocks,developers)
<h1 align="center">goradio</h1>

<h3 align="center">The simple cli radio written in Golang<h3><h2>Features</h2>

- Fast CLI
- Using mpv to play radio
- Wide variety of radio stations
- Cross-Platform - Linux, macOS, Windows
- Colorized TUI

<h2>Installation</h2>

<h3>Golang (Windows, Linux, MacOS)</h3>

Install using [Golang Packages](https://pkg.go.dev/github.com/neptunsk1y/goradio)

```shell
go install github.com/neptunsk1y/goradio@latest
```

This script will automatically detect OS & Distro and use the best option available.


<h3> From source </h3>

Clone the repo
```shell
git clone https://github.com/neptunsk1y/goradio.git
cd goradio
```

GNU Make **(Recommended)**
```shell
make setup # if you want to compile and install goradio to path
```

<details>
<summary>If you don't have GNU Make use this</summary>


```shell
# To build
go build

# To install
go install
```

</details>

<h2>Usage</h2>

<h3>Radio</h3>

To run: `goradio radio`

![Radio](https://github.com/neptunsk1y/goradio/blob/main/assets/goradio.gif?raw=true)

<details>
<summary>Keybinds</summary>

| Bind         | Description       |
|--------------|-------------------|
| <kbd>↑</kbd> | Prev/Up station   |
| <kbd>↓</kbd> | Next/Down station |
| <kbd>→</kbd> | Next page         |
| <kbd>←</kbd> | Prev page         | 
| <kbd>/</kbd> | Search            |
</details>


<h3>Other</h3>

See `goradio help` for more information

<h2> Built With </h2>

* [Cobra](https://cobra.dev/) - The modern CLI framework used

<h2> Contributing </h2>

Please read [CONTRIBUTING.md](https://github.com/neptunsk1y/goradio/blob/main/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

<h2> Authors </h2>

* **Mikhail Chikankov** - *Creator project* - [neptunsk1y](https://github.com/neptunsk1y)


<h2>License</h2>

Sample and its code provided under MIT license, please see [LICENSE](/LICENSE). All third-party source code provided
under their own respective and MIT-compatible Open Source licenses.

Copyright (C) 2023, Mikhail Chikankov


<h2> Acknowledgments </h2>

* Hat tip to anyone whose code was used
* Inspiration
* etc

# Clarity by Example

Content and build toolchain for [Clarity by Example](https://claritybyexample.com),
a site that teaches Clarity via annotated example smart-contracts. Based on [Go by Example](https://github.com/mmcgrana/gobyexample)

### Overview

The Clarity by Example site is built by extracting code and
comments from source files in `examples` and rendering
them via the `templates` into a static `public`
directory. The programs implementing this build process
are in `tools`, along with dependencies specified in
the `go.mod`file.

The built `public` directory can be served by any
static content system. 

### Building

To build the site you'll need Go installed. Run:

```console
$ tools/build
```

To build continuously in a loop:

```console
$ tools/build-loop
```

To see the site locally:

```
$ tools/serve
```

and open `http://127.0.0.1:8000/` in your browser.

### Thanks

Thanks to [Mark McGranaghan](https://markmcgranaghan.com) and [Eli Bendersky](https://eli.thegreenplace.net) for building [GoByExample](https://github.com/mmcgrana/gobyexample) the inspiration for this project.

### License 
This work is copyright Linden and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

#### Go By Example
This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](http://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).







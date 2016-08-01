# UniDoc

[UniDoc](http://unidoc.io) is a fast and powerful open source library for document manipulation starting off as a PDF
toolkit. This is a commercial library written and supported by the owners
of the FoxyUtils.com website.

This library is used to power many of the services offered by the FoxyUtils.com website. The goal is to extend it to
eventually support all of the offered services.

[![wercker status](https://app.wercker.com/status/22b50db125a6d376080f3f0c80d085fa/s/master "wercker status")](https://app.wercker.com/project/bykey/22b50db125a6d376080f3f0c80d085fa)
[![GoDoc](https://godoc.org/github.com/unidoc/unidoc?status.svg)](https://godoc.org/github.com/unidoc/unidoc)

## Getting the code

Open source users can use the master branch.
Commercial users get a special URL with their customer id. Only the commercial URLs are eligible for commercial support.

## Installation
~~~
go get github.com/unidoc/unidoc
~~~

## Overview

 * Many [features](http://unidoc.io/features) with documented examples.
 * Self contained with no external dependencies
 * Developer friendly

## Examples

See the [unidoc-examples](https://github.com/unidoc/unidoc-examples/tree/master) folder.

## Roadmap

The following features are on the roadmap, these are all subjects to change.

 * Compress PDF
 * Create PDF (high level API)
 * Fill out Forms
 * Create Forms
 * Bindings for Python (and C#/Java if there is interest)
 * Create Doc and DocX files
 * Convert PDF to Word
 * OCR Engine
 * And many more...

## Copying/License

UniDoc is licensed as [AGPL][agpl] software (with extra terms as specified in our license).

AGPL is a free / open source software license.

This doesn't mean the software is gratis!

Buying a license is mandatory as soon as you develop commercial activities
distributing the UniDoc software inside your product or deploying it on a network
without disclosing the source code of your own applications under the AGPL license.
These activities include:

 * developing commercial applications
 * offering paid services to customers as an application service provider
 * serving PDFs on the fly in the cloud or in a web application
 * shipping UniDoc with a closed source product

Contact sales for more info: sales@unidoc.io.

## Contributing

Contributors need to approve the [Contributor License Agreement](https://docs.google.com/a/owlglobal.io/forms/d/1PfTjEAi67-x0JOTU45SDonJnWy1fWB_J1aopGss34bY/viewform) before any code will be reviewed. Preferably add a test case to make sure there is no regression and that the new behaviour is as expected.

## Support

Open source users can create a GitHub issue and we will look at it. Commercial users can either create a GitHub issue and also email us at support@unidoc.io and we will assist them directly.

## Stay up to date

* Follow us on [twitter](https://twitter.com/unidoclib)
* Sign-up for our [newsletter](http://eepurl.com/b9Idt9)

[agpl]: LICENSE.md
[contributing]: CONTRIBUTING.md
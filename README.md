xorm inflector mapper
=====================
[![Vexor status](https://ci.vexor.io/projects/b4ca5d2b-6477-4056-9ef6-9056b68e3b65/status.svg)](https://ci.vexor.io/ui/projects/b4ca5d2b-6477-4056-9ef6-9056b68e3b65/builds)

Package implements xorm IMapper that pluralizes table names

Installation
============

Just use go get.

```golang
go get github.com/humaniq/xorm-inflector-mapper
```

And then just import the package into your own code.

```golang
import (
    "github.com/humaniq/xorm-inflector-mapper"
)
```

Usage
=====

You should use `InflectorMapper` with other [xorm](github.com/go-xorm/xorm) mappers:

```golang
engine.SetTableMapper(mapper.NewInflectorMapper(SnakeMapper{}))
```

Pull requests policy
====================

tl;dr. Contributions are welcome.

xorm inflector mapper
=====================
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

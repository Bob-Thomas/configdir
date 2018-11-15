configdir for Go
================

Greatly reduced version of the [library from shibukawa](https://github.com/shibukawa/configdir).

Usage
=====

```
dir := configdir.SettingsDir("example-vendor", "example-app")
```
```
dir := configdir.CacheDir("example-vendor", "example-app")
```
```
dir := configdir.SystemSettingsDir("example-vendor", "example-app")
```

Documentation
-------------

https://godoc.org/github.com/tfriedel6/configdir

License
-------

MIT


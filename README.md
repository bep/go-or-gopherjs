# go-or-gopherjs

A CLI wrapper invoking `go` or `gopherjs` based on an environment variable.

Simplest way of installing:

`go get github.com/bep/go-or-gopherjs` and make sure that `$GOPATH/bin` is on your `PATH`.

# Use

I use [LiteIDE](https://github.com/visualfc/liteide) for all my Go development, so the steps provided here works for that application. This wrapper may not even be needed in other IDEs. Also note that the paths provided are for `macOS` and will be different on other platforms.

Edit `/Applications/LiteIDE.app/Contents/Resources/litebuild/gosrc.xml`:

```
<config id="Go" name="GO" value="go"/>
```

To:

```
<config id="Go" name="GO" value="go-or-gopherjs"/>
```

Then, in `/Applications/LiteIDE.app/Contents/Resources/liteenv`:

```
cp system.env js.env
```

Add:

```
MYGOARCH=js
```

TODO: GOARCH vs MYGOARCH, see https://github.com/bep/go-or-gopherjs/issues/1

Now you can just switch to the `js` environment inside LiteIDE when you are doing GopherJS development.



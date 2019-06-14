# gopls modifying mod

When executing a command `gopls` in a root folder, with `sub/go.mod` in a subfolder, then `gopls` ends up adding dependencies to `go.mod`, which should not be there.

In this example:

```
gopls format ./sub/main.go
```

Creates a diff looking like:

```
diff --git a/gopls-modifying-parent-mod/go.mod b/gopls-modifying-parent-mod/go.mod
index 37ecab5..2bb0627 100644
--- a/gopls-modifying-parent-mod/go.mod
+++ b/gopls-modifying-parent-mod/go.mod
@@ -1,3 +1,5 @@
 module example

 go 1.12
+
+require github.com/kylelemons/godebug v1.1.0 // indirect
```

Internally it seems to be invoking:

```
go list -e -json -compiled=true -test=true -export=false -deps=true -find=false -- fmt github.com/kylelemons/godebug/diff
```

Which modifies `go.mod` and `go.sum`.
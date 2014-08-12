gocheck-samepath
================

A gocheck checker which tests whether Windows-esque backslash-separated paths are the same with slash-separated paths.

* Usage examples may be found in the _test.go file.

* Notes:
  * The checker will work for paths featuring both slashes and backslashes regardless of the OS it is currently running on.
  * The checker is capable of handling Windows shortened paths, junction points and Linux *hard* links (symlinks in linux are handled differently than the way this checker does it).
  * This checker was meant for integration(aka copy-pasting) into the juju/testing/checkers package and thus relies on the stringOrStringer function from it; in the version off this exact repo, stringOrStringer is imported(to do this you'll have to export if first).

Extra credit to [Bogdan Teleaga](https://github.com/bogdanteleaga) for his
contributions to the checker.

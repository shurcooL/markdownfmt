Test cases
==========

Definition
----------

Test cases can be defined by creating a file inside `default` (for the default
test suite) or by creating a new directory in `testdata`. This new directory
should contain a set of input files (`*.in.md`) and matching output files
(`*.out.md`) as well as a single `options.json` that defines the set of options
to invoke the markdown formatter with. Options are defined as JSON objects that
have a key/value for each of the markdown formatter options. For example, to
define the `Terminal` option, you would define `options.json` to be

```
{
  "terminal": true
}
```

Updating test cases
-------------------

Since it is really annoying to write out test case expected outputs by hand you
can use the `-update_goldens` flag when invoking `go test` to have it
automatically write out the `.out.md` file expected by the failed testcase.

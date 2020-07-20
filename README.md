## k8syaml2json

Usage: k8syaml2json < $inputYamlFile > $newlineSeparatedOutputJsonFile

k8syaml2json converts the yaml document or documents input via stdin into json,
and outputs them on stdout.
It follows kubernetes semantics for both yaml and json.

The output is simple newline separated json blobs, one per input yaml document.

Multiple yaml documents may be separated by '---'.

### Why use this?

Some tooling is better at dealing with json than yaml (such as nix). This tool
is a helper for those programs.

In addition, various languages that aren't go are really bad at deserializing
k8s yaml, and are much better at deserializing k8s json.

This can also help act as a bridge for those languages.

### What about CRDs?

They probably don't work. PRs welcome.

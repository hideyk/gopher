## 3.1 Aliaing
Packages can be prefixed with an alias which is useful when we experience naming collisions, or wish to work with package contents using a more semantic name

Aliasing should be used judiciously to add clarity and not mischeviously!

## 3.2 Dot import prefix
We can use a dot prefix to promote the functionality of an imported package to the current package namespace.

Rarely see in practice as it mainly distracts from clarity.

## 3.3 Init functions
`init()` functions execute before `main()` but after package level variables are evaluated. Their use should be constrained to reading environment variables or a config file.
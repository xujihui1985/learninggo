### Mechanics

In Go, all packages are "first class", and the only hierarchy is what you define in the source tree
for your project

Two packages can't cross import each other. imports are a one way steet

### Design

#### to be purposeful, packages must provide, not contain
#### to be usable, packages must be designed with the user as their focus
#### packages must be intuitive and simple to use
#### Packages must protect the users application from cascading changes
#### packages must reduce taking on opinions when it's reasonable and practical
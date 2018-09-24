### rule

#### cmd is where binary live, we should put all related binary in one library
#### internal is the pkg which can not be seen by other project, they are reusable across multiple binaries
#### package of same level under internal can not import each other, they must have a separate, portable API
>what happened if orders and attachment need to use the same database tables or same datastruct?, they have to have their own version of 
struct, they need to maintain a level of decoupling

#### in this case `orders` can import `customers` and **leverage it's API**
#### when you import a package only because you need a type, that is a flag, we want an import because we need that behavior, not the type
#### any internal package can import platform, you can call it common/util

### when you want to import package in the same level, think about it
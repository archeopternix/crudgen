# CRUDgen
Generator for a web based CRUD application/API with selectable frontends and backends. 
CRUDgen uses an AST (abstract syntax tree) that is backed up by a configuration files in YAML. 
You will be provided with an interactive command line interface that helps you building up the 
model which is the foundation for the code generation.

## Application
A full (web based) CRUD application is generated using templates 'go/template'.
Such an application consists of 4 base components where frontends and backends 
could be easily exchanged. It is even possible to write own frontend or backend adapters

### Core:
Configuration, interfaces and code that glues the frontend, backend and model together.

### Models:
crudgen is a generator that uses go/template package to generate models,
base logic 

### Frontends (could be an API):
* webpages/routers for echo V4 framework (https://echo.labstack.com/)

### Backends:
* in-memory databases
* SQL databases

## Command line interface CLI:
Command line interface

### Initialisation
Basic setup of the AST configuration in the target directory. 
Configuration files will be created with default data set.

    Usage:
      crudgen init [flags]

    Flags:
       -h, --help              help for init
       -n, --name string       Name of the application
           --pkg-name string   Package name of the root package (e.g. github.com/abc)

### Adding an entity
An entity will be added to the configuration. The default type is a
normal 'entity' that holds fields, it is necessary to create fields and add 
them to the entity configuration.

A special entity type is 'lookup' which could populate drop down fields.

    Usage:
      crudgen add entity [flags]

    Flags:
      -h, --help          help for entity
      -n, --name string   Name of the entity
      -t, --type string   Type of the entity to be created (default or lookup (default "default")

### Adding a relation
The relation will be added to the configuration. You can choose as 
relation type (e.g.) onetomany. As a flag source and target have to be submitted as 
the both entitites that are in a relation to each other

    Usage:
       crudgen add relation [flags]

    Flags:
      -h, --help            help for relation
      -s  --source string   Name of the source (e.g. 1..) entity
      -t  --target string   Name of the target (e.g. ..n) entity
          --type string     Type of relation (1..n = onetomany) (default "onetomany")

### Adding a text field (text)
Adds a text field to an entity where you can set if the field is --required 
or used as a --label in drop down select boxes and define the maximum length. 
Length=-1 means no restriction

    Usage:
      crudgen add text [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
          --label           This field will be used as a label for drop down fields (to activate: --label)
      -l, --length int      Maximum text length (-1 .. means no restriction) (default -1)
      -n, --name string     Name of the field
          --required        Content for field is required to be accepted (to activate: --required)

### Adding a longtext field (longtext)
Adds a longtext field to an entity where you can set if the field is --required 
and define the maximum length. Length=-1 means no restriction

    Usage:
      crudgen add longtext [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
      -l, --length int      Maximum text length (-1 .. means no restriction) (default -1)
      -n, --name string     Name of the field
          --required        Content for field is required to be accepted (to activate: --required)

### Adding a password field (password)
Adds a password field to an entity where you can set if the field is --required 
and define the maximum length. Length=-1 means no restriction

    Usage:
      crudgen add password [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
      -l, --length int      Maximum text length (-1 .. means no restriction) (default -1)
      -n, --name string     Name of the field
          --required        Content for field is required to be accepted (to activate: --required)

### Adding a phone field (phone)
Adds a phone field to an entity where you can set if the field is --required 
and define the maximum length. Length=-1 means no restriction

    Usage:
      crudgen add phone [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
      -l, --length int      Maximum text length (-1 .. means no restriction) (default -1)
      -n, --name string     Name of the field
          --required        Content for field is required to be accepted (to activate: --required)


### Adding a e-mail field (email)
Adds a e-mail field to an entity where you can set if the field is --required 
and define the maximum length. Length=-1 means no restriction

    Usage:
      crudgen add email [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
      -l, --length int      Maximum text length (-1 .. means no restriction) (default -1)
      -n, --name string     Name of the field
          --required        Content for field is required to be accepted (to activate: --required)

### Adding a boolean field (boolean)
Adds a boolean (true/false) field to an entity.

    Usage:
      crudgen add boolean [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
      -n, --name string     Name of the field
    
add integerfield

add numberfield

add lookupfield




add timefield


### configure
- configure view / repository -kind/ -help

### generate
- generate all / model / view / repository

### run
- run

## AST tree

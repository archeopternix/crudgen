# CRUDgen
Generator for a web based CRUD application/API with selectable frontends and backends. 

The arget is to create a full (web based) CRUD application that is generated using templates 'go/template'.
Such an application consists of frontends (view components or API's) and backends (repositories like MOCK or SQL databases).
These modules are organized in a different [CRUDgen modules](https://github.com/archeopternix/crudgen-modules) repository that will be loaded dynamically upon code generation
It is even possible to write own frontend or backend adapters

CRUDgen uses an [AST](https://github.com/archeopternix/crudgen/ast)  (abstract syntax tree) that is backed up by a configuration files in YAML. 
You will be provided with an interactive command line interface that helps you building up the 
model which is the foundation for the code generation.

CRUDgen provides
* Easy subcommand-based CLIs: crudgen run, crudgen add entity, etc.
* Fully POSIX-compliant flags (including short & long versions)
* Nested subcommands
* Easy generation of applications & commands with cobra init appname & cobra add cmdname
* Automatic help flag recognition of -h, --help, etc.
* Pluggable modules stored in a separate repository


# Installing
Using CRUDgen is easy. First, use `go get` to install the latest version
of the library. This command will install the `crudgen` generator executable
along with the library and its dependencies:

    go get github.com/archeopternix/crudgen
    
    

## Command line interface CLI:

CRUDgen CLI is built on a structure of commands, arguments & flags.

**Commands** represent actions, **Args** are things and **Flags** are modifiers for those actions.

The best applications read like sentences when used, and as a result, users
intuitively know how to interact with them.

The pattern to follow is
`crudgen VERB NOUN --ADJECTIVE.`
    or
`crudgen COMMAND ARG --FLAG`

### Initialisation
Basic setup of the AST configuration in the target directory. 
Configuration files will be created with default data set.

    Usage:
      crudgen init [flags]

    Flags:
       -h, --help              help for init
       -n, --name string       name of the application
           --pkg-name string   package name of the root package (e.g. github.com/abc)

### Adding an entity
An entity will be added to the configuration. The default type is a
normal 'entity' that holds fields, it is necessary to create fields and add 
them to the entity configuration.

    Usage:
      crudgen add entity [flags]

    Flags:
      -h, --help          help for entity
      -n, --name string   name of the entity
      -t, --type string   type of the entity to be created (default "default")

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
          --columns int     Columns for textfield (default = 80) (default 80)
          --rows int        Rows for textfield (default = 4 (default 4)

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
    
### Adding an integer field (integer)
Adds an integer field to an entity where you can set the 'min', 'max' value 
that is allowed to enter. The standard 'step' between values is 1 (means integer) but this can 
be changed by setting the 'step' flag

    Usage:
      crudgen add integer [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
      -n, --name string     Name of the field
          --max int         Minimum value for field ((default no maximum (default 9223372036854775807)
          --min int         Minimum value for field (default no minimum) (default -9223372036854775808)
          --step int        Step between values (step = 1 for integer) (default 1)

### Adding a number field (number)
Adds a number field to an entity. Numbers are any floating point values
    Usage:
      crudgen add number [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
      -n, --name string     Name of the field
    
### Adding a lookup field (lookup)
Adds a lookup field to an entity. The name of the lookupfield needs to 
be the name of the corresponding entity

    Usage:
      crudgen add lookup [flags]

    Flags:
      -e, --entity string   Entity where the field will be added
      -h, --help            help for fieldtext
      -n, --name string     Name of the field


### configure
- configure view / repository -kind/ -help

### generate
- generate all / model / view / repository

### run
- run

## AST tree

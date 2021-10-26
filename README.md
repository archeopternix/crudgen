# CRUDgen
Generator for a web based CRUD application/API with selectable frontends and backends. 
CRUDgen uses an AST tree that will build up based on configuration files in YAML. 
You will be provided with interactive shell commends that helps you building up the 
application stack

## Application
a full (web based) CRUD application is generated using templates 'go/template'.
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

## Command line:
Command line interface

### add
- add entity / relation -name / -help

### configure
- configure view / repository -kind/ -help

### generate
- generate all / model / view / repository

### run
- run

## AST tree

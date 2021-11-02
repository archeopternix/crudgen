# AST (abstract syntax tree)

AST is a tree representation of the abstract syntactic structure of text (often source code) written in a formal language. Each node of the tree denotes a construct occurring in the text. The syntax is "abstract" in the sense that it does not represent every detail appearing in the real syntax, but rather just the structural or content-related details. 

## Application
Application holds all information and configuration for the AST and consists
of 
* Entitites
* Relations 
* Configuration necessary for template generation

## Entity 
Entity could be seen as a database table and holds the field definitions

## Field
Field is the definition for every single attribute within an entity.
Right now these field types are supported:
* text
* password
* integer
* number
* boolean
* email
* tel
* longtext
* time
* lookup

## Relation
Relation holds the definition for parent - child relationships.
When parsed by Application additional fields will be added to the child and parent
entities
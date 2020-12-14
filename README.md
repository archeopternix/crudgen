# crudgen
Generator for a web based CRUD application with echo frontend and pluggable storage backends

crudgen is a generator that uses go/template package to generate 
the base logic, structs (objects), and webpages/routers for echo V4 framework (https://echo.labstack.com/).
Application data can be stored in diffeent repositories and databases

Commands:

create:
Creates a new application with a default landingpage
    create NAME
    


    server --port:8080
    add entity NAME
    add field NAME -entity:ENAME --type:TYPE

# ZONA DNS cloudnativerioja.com

## Descripción

Este repositorio contiene el código necesario para desplegar una infraestructura en Cloudflare utilizando Pulumi en Golang.

## Objetivo

El objetivo de este proyecto es gestionar los records de la zona DNS cloudnativerioja.com por código, de forma que se pueda desplegar la infraestructura de forma automática y replicable.

## Uso

Para modificar, añadir o eliminar records de la zona DNS:

- Se crea una nueva rama a partir de `master`
- Se debe modificar el fichero `main.go` para crear/modificar/eliminar los records necesarios
- Se crea una Pull Request para que se revise el código en la pipeline creada para ello
- Se añade Code Reviewers al grupo Collaborators para que al menos una persona revise el código y lo válide
- Se mergea la Pull Request a `master` y las pipelines de despliegue y release crearán/modificarán/eliminarán los records necesarios en la zona DNS

## Eliminación de todos los records (probablemente no se necesite nunca)

Para la eliminación de todos los records en cloudflare:

- Crear tag con el nombre `*destroy*`, esto disparará la pipeline para eliminar todos los records de la zona DNS

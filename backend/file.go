package main

import (
	"os"
	"path/filepath"
)

// la funcion devuelve true si se crea correctamente el fichero
func create_file(nombreFichero string, pathFichero string) bool {
	// esta será la ruta completa
	fullPath := filepath.Join(pathFichero, nombreFichero)

	// Crear fichero
	file, err := os.Create(fullPath)
	if err != nil {
		return false
	}
	defer file.Close()
	return true
}

func mod_file(nombreFichero string, pathFichero string, contenFichero string) bool {

	fullPath := filepath.Join(pathFichero, nombreFichero)
	//abrimos el fichero
	//O_WRONLY --> para abrir en modo escritura
	//O_TRUNC --> para que guarde el contenido sin concadenar con lo que hay
	//O_CREATE --> para que si no existe el fichero lo cree
	//0777 --> permisos del fichero
	ficheroAbierto, log := os.OpenFile(fullPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if log != nil {
		return false
	}
	// Escribe en el archivo
	_, log = ficheroAbierto.WriteString(contenFichero)
	if log != nil {
		return false
	}

	return true
}

func existeFichero(nombreFichero string, pathFichero string) bool {
	res := true
	_, log := os.Stat(filepath.Join(pathFichero, nombreFichero))
	if os.IsNotExist(log) {
		res = false
	}
	return res
}


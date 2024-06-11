package main

import (
	"os"
	"path/filepath"
)

// la funcion devuelve true si se crea correctamente el fichero
func create_file(nombreFichero string, pathFichero string) bool {
	// esta será la ruta completa
	res := true
	fullPath := filepath.Join(pathFichero, nombreFichero)

	// Crear fichero
	file, err := os.Create(fullPath)
	if err != nil {
		res = false
	}
	defer file.Close()
	return res
}

// en caso de mandar modificar un file que no existe lo crea, puede que con esta podamos prescindir de create_file()
// ? utilizar create_file para crear ficheros vacios o usar solo esta
func mod_file(nombreFichero string, pathFichero string, contenFichero string) bool {
	res := true
	fullPath := filepath.Join(pathFichero, nombreFichero)
	//abrimos el fichero
	//O_WRONLY --> para abrir en modo escritura
	//O_TRUNC --> para que guarde el contenido sin concadenar con lo que hay
	//O_CREATE --> para que si no existe el fichero lo cree
	//0777 --> permisos del fichero
	ficheroAbierto, log := os.OpenFile(fullPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if log != nil {
		res = false
	}
	defer ficheroAbierto.Close() //Esto es necesario para evitar desbordamiento de datos, lo que hace es que se cierre el fichero al finalizar al funcion

	// Escribe en el archivo
	_, log = ficheroAbierto.WriteString(contenFichero)
	if log != nil {
		res = false
	}

	return res
}

func existeFichero(nombreFichero string, pathFichero string) bool {
	res := true
	_, log := os.Stat(filepath.Join(pathFichero, nombreFichero))
	if os.IsNotExist(log) {
		res = false
	}
	return res
}

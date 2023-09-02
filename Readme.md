# FileManager - Librería de Gestión de Archivos en Go

FileManager es una librería en Go diseñada para simplificar la gestión de archivos. Su primera función, `Copy`, permite copiar archivos de manera sencilla y eficiente.

## Instalación

Para usar FileManager en tu proyecto Go, primero debes asegurarte de tener Go instalado. Luego, puedes instalar FileManager utilizando el siguiente comando:

```shell
go get github.com/Augusto-p/FileManager
```


## Funciones

### Copy(sourcePath string, destinationPath string) error

La función `Copy` permite copiar un archivo desde una ubicación de origen a una ubicación de destino especificadas. Esta función toma dos argumentos: `sourcePath`, que es la ruta del archivo de origen que se copiará, y `destinationPath`, que es la ruta donde se copiará el archivo.

Ejemplo de uso:

```go
sourcePath := "archivo_origen.txt"
destinationPath := "archivo_destino.txt"

err := FileManager.Copy(sourcePath, destinationPath)
if err != nil {
    fmt.Println("Error al copiar el archivo:", err)
    return
}

fmt.Println("Archivo copiado exitosamente.")
```

Asegúrate de importar `github.com/Augusto-p/FileManager` en tu proyecto y reemplazar `archivo_origen.txt` y `archivo_destino.txt` con las rutas de tus archivos reales.

## Contribuciones

Si deseas contribuir a FileManager, ¡estamos abiertos a colaboraciones! Siéntete libre de crear problemas (issues) o enviar solicitudes de extracción (pull requests) en [el repositorio de GitHub](https://github.com/Augusto-p/FileManager).

## Licencia

Este proyecto se distribuye bajo la licencia MIT. Consulta el archivo [LICENSE](LICENSE) para obtener más detalles.

<!-- ## Contacto

Si tienes alguna pregunta o comentario sobre FileManager, no dudes en ponerte en contacto con nosotros en [correo@example.com](mailto:correo@example.com). -->

¡Gracias por usar FileManager!
```

Asegúrate de personalizar este `README.md` con la información real de tu librería y proporcionar ejemplos de uso más detallados y ejemplos adicionales si es necesario. También, considera incluir ejemplos de manejo de errores y casos de uso más avanzados si tu librería los admite.
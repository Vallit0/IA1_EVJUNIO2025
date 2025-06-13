#!/bin/bash

# Reemplaza esto con tu CARNET real
CARNET="X"
REPO="IA1_1S2025_$CARNET"

# Crear estructura de carpetas
mkdir -p $REPO/backend/prolog $REPO/docs
cd $REPO

# Archivos esenciales
echo "# Proyecto UniMatch – Backend" > README.md
echo "MIT License" > LICENSE
touch docs/manual_usuario.md docs/manual_tecnico.md

# Backend base en Go
cat <<EOF > backend/main.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Servidor UniMatch en ejecución.")
    })

    fmt.Println("Servidor escuchando en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
EOF

cat <<EOF > backend/go.mod
module unmatch/backend

go 1.20
EOF

# Archivo base para reglas en Prolog (Tau-Prolog)
touch backend/prolog/conocimiento.pl

# Inicializar Git y primer commit
git init
git add .
git commit -m "Inicialización del backend UniMatch"

echo "✅ Backend creado en ./$REPO"
echo "📌 Recuerda subir esto a GitHub como repositorio privado con nombre exacto:"
echo "   ➤ IA1_1S2025_$CARNET"
echo "🔗 Y agregar como colaborador a: vallit0"

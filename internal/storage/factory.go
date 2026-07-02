package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"atletismo-api/internal/models"
	//ciclosModels "atletismo-api/internal/models"
	//evaluacionesModels "atletismo-api/internal/models"
	//microciclosModels "atletismo-api/internal/models"
	//usuarioModels "atletismo-api/internal/models"
	//viviendaModels "atletismo-api/internal/models"
)

// Recursos agrupa todo lo que la capa de almacenamiento expone a la aplicacion.
type Recursos struct {
	Almacen      Almacen
	Usuarios     UserRepository
	BackendUsado string
	Cerrar       func() error
}

// Inicializar centraliza TODO el plumbing de almacenamiento (patron Factory).

// Inicializar centraliza TODO el plumbing de almacenamiento (patron Factory).
func Inicializar(rutaDB, backend string) (*Recursos, error) {
	// 1. GORM es el DUENO DEL ESQUEMA: abre, migra y siembra.
	gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("abrir GORM: %w", err)
	}
	if err := gdb.AutoMigrate(
		&models.CicloEntrenamiento{},
		&models.EvaluacionCiclo{},
		&models.Microciclo{},
		&models.Usuario{},
	); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}
	almacenGorm := NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend de productos/categorias.
	//    El backend sqlc esta generado para SQLite (sus queries son de SQLite),
	//    por eso solo aplica cuando el driver es sqlite; con postgres se usa GORM.
	var almacen Almacen
	var sdb *sql.DB
	backendUsado := "gorm"
	switch backend {
	case "sqlc":
		sdb, err = sql.Open("sqlite", rutaDB)
		if err != nil {
			return nil, fmt.Errorf("abrir sql.DB para sqlc: %w", err)
		}
		almacen = NuevoAlmacenSQLC(sdb)
		backendUsado = "sqlc"
	default:
		almacen = almacenGorm
	}

	// 3. Usuarios viven SIEMPRE en GORM (decision tomada en S10).
	usuarios := NewUsuarioRepository(gdb)

	// 4. Cierre ordenado: primero la conexion sql.DB de sqlc (si existe), luego GORM.
	cerrar := func() error {
		if sdb != nil {
			if err := sdb.Close(); err != nil {
				return err
			}
		}
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
		Almacen:      almacen,
		Usuarios:     usuarios,
		BackendUsado: backendUsado,
		Cerrar:       cerrar,
	}, nil
}

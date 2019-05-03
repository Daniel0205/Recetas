package main


import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
    "database/sql"
	"log"
	"fmt"
    _ "github.com/lib/pq"
)

type Receta struct {
	nombre string `json:"nombre"`
	preparacion string `json:"preparacion"`
	ingredientes []Ingrediente `json:"ingredientes"`
}

type Ingrediente struct {
	
	nombre string `json:"nombre,omitempty"`
	unidad string `json:"unidad,omitempty"`
	cantidad float32  `json:"cantidad,omitempty"`
}
  


// EndPoints
func crearReceta(w http.ResponseWriter, req *http.Request){

}

func consultarReceta(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)

	var receta Receta;
	var ingrediente Ingrediente;	 
	// Connect to the "bank" database.
	db, err := sql.Open("postgres", "postgresql://chef@localhost:26257/recetas?sslmode=disable")
	 
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	q := "SELECT * FROM receta NATURAL JOIN ingredientes WHERE nombre='"+params["nombre"]+"';"
	
	// Print out the balances.
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	
	defer rows.Close()
	
	fmt.Print(q)
	fmt.Print("\n")


    for rows.Next() {
		var nombre,preparacion,nomIngrediente,unidad string 
		var cantidad float32
        if err := rows.Scan(&nombre,&preparacion,&nomIngrediente,&cantidad,&unidad); err != nil {
            log.Fatal(err)
		}
		receta.nombre = nombre
		receta.preparacion = preparacion
		ingrediente.nombre = nomIngrediente
		ingrediente.unidad = unidad
		ingrediente.cantidad = cantidad
		receta.ingredientes = append(receta.ingredientes,ingrediente)	
		
	}
	fmt.Print(receta)
	fmt.Print("\n")
	json.NewEncoder(w).Encode(receta)
}

func actualizarReceta(w http.ResponseWriter, req *http.Request){


}

func borrarReceta(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	
	// Connect to the "bank" database.
	db, err := sql.Open("postgres", "postgresql://chef@localhost:26257/recetas?sslmode=disable")
	
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	q := "DELETE FROM ingredientes WHERE nombre='"+params["nombre"]+"';"+
	     "DELETE FROM receta WHERE nombre='"+params["nombre"]+"';"
	
	_, err = db.Exec(q)
	if err != nil {
	  panic(err)
	}
}

func main() {
/*	 
	// Connect to the "bank" database.
	db, err := sql.Open("postgres", "postgresql://chef@localhost:26257/recetas?sslmode=disable")
	 
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	 // Create the "receta" table.
	 if _, err := db.Exec(
		 "CREATE TABLE IF NOT EXISTS receta (nombre TEXT PRIMARY KEY,preparacion TEXT NOT NULL);"); err != nil {
		 log.Fatal(err)
	 }
 
	 // Insert two rows into the "accounts" table.
	 if _, err := db.Exec(
		 "CREATE TABLE IF NOT EXISTS ingredientes(nombre_ingrediente TEXT NOT NULL,cantidad REAL NOT NULL, "+
			"unidad_medida varchar(5) NOT NULL,nombre TEXT,FOREIGN KEY (nombre) REFERENCES receta(nombre));"); err != nil {
		 log.Fatal(err)
	 }

	// Insert two rows into the "accounts" table.
	if _, err := db.Exec(
		"INSERT INTO receta (nombre, preparacion) VALUES ('milo frio', 'revolver todo'), ('huevo frito', 'batir')"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		"INSERT INTO ingredientes (nombre,nombre_ingrediente,cantidad ,unidad_medida) VALUES "+
		"('milo frio','leche',1.5,'Lt'), ('milo frio', 'milo',1, 'gr')"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		"INSERT INTO ingrediente (nombre,nombre_ingrediente,cantidad ,unidad_medida) VALUES "+
		"('huevo frito','aceite',0.2,'mLt'), ('huevo frito','huevo','1', 'Unid')"); err != nil {
		log.Fatal(err)
	}

*/
  router := mux.NewRouter()

  // endpoints
  router.HandleFunc("/receta", crearReceta).Methods("POST")
  router.HandleFunc("/receta/{nombre}", consultarReceta).Methods("GET")
  router.HandleFunc("/actualizar", actualizarReceta).Methods("POST")
  router.HandleFunc("/receta/{nombre}", borrarReceta).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":3000", router))
}
package main


import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
  "database/sql"
	"log"
	"fmt"
    _ "github.com/lib/pq"
)

type Receta struct {
	Nombre string `json:"nombre,omitempty"`
	Preparacion string `json:"preparacion,omitempty"`
	Ingredientes []Ingrediente `json:"ingredientes,omitempty"`
}

type Ingrediente struct {
	
	Nombre string `json:"nombre,omitempty"`
	Unidad string `json:"unidad,omitempty"`
	Cantidad float32  `json:"cantidad,omitempty"`
}
	

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}


// EndPoints
func crearReceta(w http.ResponseWriter, req *http.Request){

	enableCors(&w)

	var recet Receta 
	_ = json.NewDecoder(req.Body).Decode(&recet)

	// Connect to the "bank" database.
	db, err := sql.Open("postgres", "postgresql://chef@localhost:26257/recetas?sslmode=disable")
	
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	q:="INSERT INTO receta (nombre, preparacion) VALUES ('"+recet.Nombre+	"', '"+recet.Preparacion+"');"
	// Insert two rows into the "accounts" table.
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}

	fmt.Print(q)

	for _, element := range recet.Ingredientes {

		q:=			"INSERT INTO ingredientes (nombre,nombre_ingrediente,cantidad ,unidad_medida) VALUES "+
		"('"+recet.Nombre+"','"+element.Nombre+"',"+fmt.Sprintf("%f", element.Cantidad)+",'"+element.Unidad+"');"
		if _, err := db.Exec(q); err != nil {
			log.Fatal(err)
		}			

		fmt.Print(q)
	}	
	
	fmt.Println(recet)

	json.NewEncoder(w).Encode(recet)
	
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
		receta.Nombre = nombre
		receta.Preparacion = preparacion
		ingrediente.Nombre = nomIngrediente
		ingrediente.Unidad = unidad
		ingrediente.Cantidad = cantidad
		receta.Ingredientes = append(receta.Ingredientes,ingrediente)	
		
	}
	fmt.Print(receta)
	fmt.Print("\n")
	json.NewEncoder(w).Encode(receta)
}



func consultarNombres(w http.ResponseWriter, req *http.Request){
	

	var nombres []string

	db, err := sql.Open("postgres", "postgresql://chef@localhost:26257/recetas?sslmode=disable")
	 if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	q := "SELECT nombre FROM receta ;"
	
	
	// Print out the balances.
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	
	defer rows.Close()
	
	fmt.Print(q)
	fmt.Print("\n")


	for rows.Next() {
		var nombre string 
		
		if err := rows.Scan(&nombre); err != nil {
				log.Fatal(err)
		}
		
		nombres=append(nombres,nombre)
	}
	fmt.Print(nombres)
	fmt.Print("\n")
	json.NewEncoder(w).Encode(nombres)
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

	q := "DELETE FROM receta WHERE nombre='"+params["nombre"]+"';"
	
	_, err = db.Exec(q)
	if err != nil {
	  panic(err)
	}
	fmt.Print(q)

	json.NewEncoder(w).Encode(true)
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
	router.HandleFunc("/receta", consultarNombres).Methods("GET")
  router.HandleFunc("/receta/{nombre}", consultarReceta).Methods("GET")
  router.HandleFunc("/actualizar", actualizarReceta).Methods("POST")
  router.HandleFunc("/receta/{nombre}", borrarReceta).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders(
		[]string{"X-Requested-With", "Content-Type", "Authorization"}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
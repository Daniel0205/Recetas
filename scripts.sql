
CREATE DATABASE IF NOT EXISTS recetas;

DROP TABLE IF EXISTS recetas.receta ;
CREATE TABLE  recetas.receta (
	nombre STRING NOT NULL,
	preparacion STRING NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (nombre ASC),
	FAMILY "primary" (nombre, preparacion)
);

DROP TABLE IF EXISTS recetas.ingredientes;
CREATE TABLE recetas.ingredientes (
	nombre_ingrediente STRING NOT NULL,
	cantidad FLOAT4 NOT NULL,
	unidad_medida VARCHAR(5) NOT NULL,
	nombre STRING NULL,
	CONSTRAINT fk_nombre FOREIGN KEY (nombre) REFERENCES receta (nombre) ON DELETE CASCADE ON UPDATE CASCADE,
	INDEX ingredientes_auto_index_fk_nombre_ref_receta (nombre ASC),
	FAMILY "primary" (nombre_ingrediente, cantidad, unidad_medida, nombre, rowid)
);

CREATE USER IF NOT EXISTS chef;
GRANT ALL ON DATABASE recetas TO chef;


INSERT INTO receta (nombre, preparacion) VALUES ('milo frio', 'revolver todo'), ('huevo frito', 'batir');

INSERT INTO ingredientes (nombre,nombre_ingrediente,cantidad ,unidad_medida) VALUES 
('milo frio','leche',1.5,'Lt'), ('milo frio', 'milo',1, 'gr');

INSERT INTO ingrediente (nombre,nombre_ingrediente,cantidad ,unidad_medida) VALUES
('huevo frito','aceite',0.2,'mLt'), ('huevo frito','huevo','1', 'Unid');
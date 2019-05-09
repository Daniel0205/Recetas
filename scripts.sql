
CREATE DATABASE IF NOT EXISTS recetas;

DROP TABLE IF EXISTS receta ;
CREATE TABLE  receta (
	nombre STRING PRIMARY KEY,
	preparacion STRING NOT NULL
);

DROP TABLE IF EXISTS ingredientes;
CREATE TABLE ingredientes (
	nombre_ingrediente STRING NOT NULL,
	cantidad FLOAT4 NOT NULL,
	unidad_medida VARCHAR(5) NOT NULL,
	nombre STRING NULL,
	CONSTRAINT fk_nombre FOREIGN KEY (nombre) REFERENCES receta (nombre) ON DELETE CASCADE ON UPDATE CASCADE,
	INDEX ingredientes_auto_index_fk_nombre_ref_receta (nombre ASC)
);

CREATE USER IF NOT EXISTS chef;
GRANT ALL ON TABLE receta TO chef;
GRANT ALL ON TABLE ingredientes TO chef;


INSERT INTO receta (nombre, preparacion) VALUES ('milo frio', 'revolver todo'), ('huevo frito', 'batir');

INSERT INTO ingredientes (nombre,nombre_ingrediente,cantidad ,unidad_medida) VALUES 
('milo frio','leche',1.5,'Lt'), ('milo frio', 'milo',1, 'gr');

INSERT INTO ingredientes (nombre,nombre_ingrediente,cantidad ,unidad_medida) VALUES
('huevo frito','aceite',0.2,'mLt'), ('huevo frito','huevo','1', 'Unid');

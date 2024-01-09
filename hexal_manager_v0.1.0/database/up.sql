-- Creación de tabla Module
CREATE TABLE IF NOT EXISTS module (
    id SERIAL PRIMARY KEY,
    
    name VARCHAR(255) NOT NULL
);

-- Creación de tabla Operations
CREATE TABLE IF NOT EXISTS operations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    module_id INTEGER REFERENCES module(id)
);

-- Creación de tabla Rol_operation
CREATE TABLE IF NOT EXISTS rol_operation (
    id SERIAL PRIMARY KEY,
    rol_id INTEGER REFERENCES rol(id),
    operation_id INTEGER REFERENCES operations(id)
);

-- Creación de tabla Rol
CREATE TABLE IF NOT EXISTS rol (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Creación de tabla User
CREATE TABLE IF NOT EXISTS user (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    rol_id INTEGER REFERENCES rol(id)
);

-- Creación de tabla Category
CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Creación de tabla Area
CREATE TABLE IF NOT EXISTS area (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INTEGER REFERENCES category(id)
);

-- Creación de tabla Proyect
CREATE TABLE IF NOT EXISTS proyect (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    area_id INTEGER REFERENCES area(id)
);

-- Creación de tabla Proyect_Participant
CREATE TABLE IF NOT EXISTS proyect_participant (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES user(id),
    proyect_id INTEGER REFERENCES proyect(id)
);

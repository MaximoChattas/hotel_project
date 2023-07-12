CREATE TABLE IF NOT EXISTS hotels (
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(300) NOT NULL,
    room_amount INT NOT NULL,
    description VARCHAR(1000),
    street_name VARCHAR(100) NOT NULL,
    street_number INT NOT NULL,
    rate DECIMAL(8,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(300) NOT NULL,
    last_name VARCHAR(300) NOT NULL,
    dni VARCHAR(8) NOT NULL,
    email VARCHAR(300) NOT NULL UNIQUE,
    password VARCHAR(300) NOT NULL,
    role VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS reservations (
	id INT PRIMARY KEY AUTO_INCREMENT,
    start_date VARCHAR(16) NOT NULL,
    end_date VARCHAR(16) NOT NULL,
    user_id INT,
    hotel_id INT,
    amount DECIMAL(8,2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

CREATE TABLE IF NOT EXISTS amenities (
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(300) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS images (
	id INT PRIMARY KEY AUTO_INCREMENT,
    path VARCHAR(300) NOT NULL,
    hotel_id INT,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)	
);

CREATE TABLE IF NOT EXISTS hotel_amenities (
    hotel_id INT,
    amenity_id INT,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id),
    FOREIGN KEY (amenity_id) REFERENCES amenities(id)
);

INSERT INTO hotels (name, room_amount, description, street_name, street_number, rate) VALUES
('Windsor Hotel', 24, 'El Windsor Hotel & Tower presenta un interior de lujo en marmol y ofrece WiFi gratuita, terraza con piscina climatizada y habitaciones con vistas a la ciudad. Se puede reservar un servicio de masajes. El gimnasio cuenta con entrenadores personales.

El Windsor Hotel & Tower esta ubicado en pleno centro de la escena cultural de Cordoba. Se halla a 200 metros de la catedral historica de la ciudad y a 300 metros del Museo Iberoamericano de Artesanias.

Las habitaciones disponen de aire acondicionado, sillas tapizadas y suelo de moqueta o parquet. Ademas, estan equipadas con mobiliario elegante de madera, minibar y TV de pantalla plana con canales por cable. Hay carta de almohadas con 5 opciones diferentes.',
'Buenos Aires', 214, 45980.00),

('Amerian Executive', 15, 'Amerian Executive Cordoba es un hotel de categoria superior situado en el centro de la ciudad de Cordoba. Su arquitectura y decoracion interior se han inspirado en el clasico estilo europeo, combinado con el estilo neoclasico creando zonas calidas y acogedoras.',
'Bv. San Juan', 137, 32905.87),

('Sol de Piedra', 12, 'Ubicado estrategicamente en el corazon de Cordoba, en nuestras instalaciones usted podra encontrar el mas moderno y innovador diseno por pisos tematicos que le permitiran conocer la belleza de la cultura local.', 'Obispo Trejo', 137, 27209),

('Orfeo Suites', 18, 'La busqueda del hotel para familias ideal en Cordoba no tiene por que ser complicada. Bienvenido a Orfeo Suites Hotel, una fantastica opcion para viajeros como tu.

Para aquellos interesados en visitar puntos de referencia conocidos durante su viaje a Cordoba, Orfeo Suites Hotel se encuentra cerca de Camino de las Altas Cumbres (2,3 km) y Templo de Cordoba La Iglesia De Jesucristo De Los Santos De Los Ultimos Dias (2,5 km).

Las habitaciones de los huespedes ofrecen servicios como aire acondicionado, minibar y frigorifico, y los huespedes pueden permanecer conectados con wifi gratuito que ofrece el hotel.

Orfeo Hotel tiene recepcion abierta 24 horas, servicio de habitaciones y espacio para guardar el equipaje para que la estancia sea mas agradable. El establecimiento tambien cuenta con piscina y desayuno incluido. Si vas a en coche a Orfeo Suites Hotel,
hay parking gratis disponible.', 'Rodriguez del Busto', 4086, 46301.99),

('Laplace Hotel', 20, 'Moderno hotel tres estrellas de 1100 mts2 emplazado en el barrio Villa Belgrano sobre la Av. Pedro Simon Laplace , una zona residencial unica por su belleza y encanto en la ciudad de Cordoba - Argentina.
Su diseno fue inspirado en la satisfaccion del visitante, equipando a sus 20 habitaciones con la ultima tecnologia.', 'Av. Simon Laplace', 5355, 41870.00),

('Howard Johnson', 25, 'Buscas donde alojarte en Cordoba? Entonces no te pierdas Howard Johnson Hotel Cordoba, un hotel de moda que te acerca a lo mejor de Cordoba.

Para que te sientas como en casa, las habitaciones en el hotel incluyen televisor de pantalla plana, minibar y aire acondicionado, y mantenerte conectado es facil, ya que hay wifi gratuito disponible.

Los huespedes tienen acceso a recepcion abierta 24 horas, servicio de habitaciones y conserje durante su estancia en Howard Johnson La Canada Hotel. Ademas, Howard Johnson La Canada Hotel ofrece piscina en la azotea y desayuno incluido,
lo que hara tu viaje a Cordoba incluso mas gratificante.', 'Pte J Figueroa Alcorta', 20, 54230.00);


INSERT INTO users (name, last_name, dni, email, password, role) VALUES
('Maximo', 'Chattas', '44347116', 'maxichattas@gmail.com', '$2a$10$EwlJ7rPRJPSpKtpsXchYoOs0.YpG7KAlfr42RmjECDFMelR9ICFQW', 'Admin'), -- Password: admin
('Yago', 'Gandara', '43299061', 'yagogandara@gmail.com', '$2a$10$vnzBIZ0rRDOWX96L/jLhguXPLYG/gAFTbsAMzB/8RtzB5VDuh0jKq', 'Customer'), -- Password: pass
('Santiago', 'Navas', '41857881', 'santinavas@gmail.com', '$2a$10$vnzBIZ0rRDOWX96L/jLhguXPLYG/gAFTbsAMzB/8RtzB5VDuh0jKq', 'Customer'), -- Password: pass
('Leonardo Tomas', 'Mendez Rodriguez', '43998614', 'leomendez@gmail.com', '$2a$10$vnzBIZ0rRDOWX96L/jLhguXPLYG/gAFTbsAMzB/8RtzB5VDuh0jKq', 'Customer'); -- Password: pass


INSERT INTO amenities (name) VALUES
('Wi-Fi'), ('Desayuno'), ('Gimnasio'), ('Estacionamiento'), ('Cochera cubierta'), ('Pileta'), ('Pileta climatizada'), ('Pool bar'), ('Sala de juegos');

INSERT INTO hotel_amenities (hotel_id, amenity_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 4), (1, 6),
(2, 1), (2, 2), (2, 3), (2, 5), (2, 7),
(3, 1), (3, 2), (3, 3),
(4, 1), (4, 2), (4, 3), (4, 4), (4, 5), (4, 6), (4, 7), (4, 8), (4, 9),
(5, 1), (5, 2), (5, 3), (5, 4), (5, 6),
(6, 1), (6, 3), (6, 4), (6, 5), (6, 7);

INSERT INTO reservations (start_date, end_date, user_id, hotel_id, amount) VALUES
('29-06-2023 15:00', '30-06-2023 11:00', 1, 1, 45980.00),
('01-07-2023 15:00', '02-07-2023 11:00', 2, 3, 27209.00),
('03-07-2023 15:00', '04-07-2023 11:00', 3, 6, 41870.00),
('05-07-2023 15:00', '06-07-2023 11:00', 4, 2, 32905.87),
('07-07-2023 15:00', '08-07-2023 11:00', 2, 5, 41870.00),
('09-07-2023 15:00', '10-07-2023 11:00', 1, 4, 45980.00),
('11-07-2023 15:00', '12-07-2023 11:00', 3, 6, 41870.00),
('13-07-2023 15:00', '14-07-2023 11:00', 4, 2, 32905.87),
('15-07-2023 15:00', '16-07-2023 11:00', 2, 5, 41870.00),
('17-07-2023 15:00', '18-07-2023 11:00', 1, 1, 45980.00),
('29-06-2023 15:00', '01-07-2023 11:00', 1, 1, 91960.00),
('03-07-2023 15:00', '07-07-2023 11:00', 2, 3, 108436.00),
('09-07-2023 15:00', '13-07-2023 11:00', 3, 6, 167480.00),
('15-07-2023 15:00', '17-07-2023 11:00', 4, 2, 65771.74),
('19-07-2023 15:00', '23-07-2023 11:00', 1, 5, 209350.00),
('25-07-2023 15:00', '27-07-2023 11:00', 2, 4, 65811.74),
('29-07-2023 15:00', '31-07-2023 11:00', 3, 6, 85924.00),
('02-08-2023 15:00', '06-08-2023 11:00', 4, 2, 131820.87),
('08-08-2023 15:00', '12-08-2023 11:00', 2, 5, 167480.00),
('14-08-2023 15:00', '16-08-2023 11:00', 1, 1, 65771.74);

INSERT INTO images (path, hotel_id) VALUES
('Images/1-1.JPG', 1),
('Images/1-2.JPG', 1),
('Images/1-3.JPG', 1),
('Images/1-4.JPG', 1),
('Images/1-5.JPG', 1),
('Images/1-6.JPG', 1),
('Images/1-7.JPG', 1),
('Images/1-8.JPG', 1),
('Images/2-1.JPG', 2),
('Images/2-2.JPG', 2),
('Images/2-3.JPG', 2),
('Images/2-4.JPG', 2),
('Images/3-1.JPG', 3),
('Images/3-2.JPG', 3),
('Images/3-3.JPG', 3),
('Images/3-4.JPG', 3),
('Images/3-5.JPG', 3),
('Images/4-1.JPG', 4),
('Images/4-2.JPG', 4),
('Images/4-3.JPG', 4),
('Images/4-4.JPG', 4),
('Images/4-5.JPG', 4),
('Images/5-1.JPG', 5),
('Images/5-2.JPG', 5),
('Images/5-3.JPG', 5),
('Images/5-4.JPG', 5),
('Images/5-5.JPG', 5),
('Images/5-6.JPG', 5),
('Images/6-1.JPG', 6),
('Images/6-2.JPG', 6),
('Images/6-3.JPG', 6),
('Images/6-4.JPG', 6),
('Images/6-5.JPG', 6);

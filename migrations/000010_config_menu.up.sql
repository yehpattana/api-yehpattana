CREATE TABLE Menu (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name VARCHAR(255),
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive')),
    side_bar BIT, -- Boolean column for side_bar
);

CREATE TABLE Sub_Menu (
    id INT IDENTITY(1,1) PRIMARY KEY,
    menu_id INT,
    name VARCHAR(255),
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive')),
    side_bar BIT, -- Boolean column for side_bar
    FOREIGN KEY (menu_id) REFERENCES Menu(id)
);

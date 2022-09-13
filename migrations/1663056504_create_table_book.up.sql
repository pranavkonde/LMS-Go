CREATE TABLE Books
(
    
  `id` VARCHAR(40) NOT NULL,
  `name` VARCHAR(20) NOT NULL,
  `author` VARCHAR(20) NOT NULL,
  `price` INT NOT NULL,
  `totalcopies` INT NOT NULL,
  `status` VARCHAR(20) NOT NULL,
  `availablecopies` INT NOT NULL,
  PRIMARY KEY (id)
);
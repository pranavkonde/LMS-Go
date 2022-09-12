CREATE TABLE Users
(
  `id` VARCHAR(10) NOT NULL,
  `first_name` VARCHAR(20) NOT NULL,
  `last_name` VARCHAR(20) NOT NULL,
  `gender` VARCHAR(10) NOT NULL,
  `age` INT NOT NULL,
  `address` VARCHAR(50) NOT NULL,
  `email` VARCHAR(30) NOT NULL,
  `password` VARCHAR(20) NOT NULL,
  `mob_no` VARCHAR(50) NOT NULL,
  `role` VARCHAR(10) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (email)
);
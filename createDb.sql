 CREATE TABLE `drivers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `license_number` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1353 DEFAULT CHARSET=utf8;

CREATE TABLE `metrics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `value` int(11) DEFAULT NULL,
  `lon` decimal(18,12) NOT NULL,
  `timestamp` datetime NOT NULL,
  `lat` decimal(18,12) NOT NULL,
  `driver_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45042 DEFAULT CHARSET=utf8;

ALTER TABLE metrics ADD CONSTRAINT fk_drivers_id FOREIGN KEY (driver_id) REFERENCES drivers(id);

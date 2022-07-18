CREATE TABLE IF NOT EXISTS todo(
	id INT auto_increment not null,
	title varchar(20) not null,
	status BOOLEAN default False,
  	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY(id)
)


/* Accounts Passwords are MD5*/
CREATE TABLE accounts(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(120),
	password VARCHAR(120)
);

INSERT INTO accounts VALUES ('admin','21232f297a57a5a743894a0e4a801fc3'); /* Username: admin Password: admin in MD5 */

/* Bot Information Table */
CREATE TABLE clients(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    guid VARCHAR(120),
    ip VARCHAR(120),
	whoami VARCHAR(120),
	os VARCHAR(120),
	installdate	VARCHAR(120),
	isadmin VARCHAR(120),
	antivirus VARCHAR(120),
	cpuinfo VARCHAR(120),
	gpuinfo VARCHAR(120),
	clientversion VARCHAR(120),
	lastcheckin VARCHAR(120),
	lastcommand VARCHAR(120)
);

/* TaskMngr */
CREATE TABLE tasks(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(120),
	guid VARCHAR(120),
	command TEXT(65533),
	method VARCHAR(120)
);

/* Commands */
CREATE TABLE command(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	command TEXT(65533),
	timeanddate VARCHAR(120)
);

/* LastC&C */
CREATE TABLE lastlogin(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	timeanddate VARCHAR(120)
);
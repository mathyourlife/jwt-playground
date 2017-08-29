-- As super user

DROP DATABASE IF EXISTS {{ .DB.Name }};
DROP USER IF EXISTS {{ .DB.Admin.Username }};
DROP ROLE IF EXISTS {{ .DB.AppUser.Username }};

CREATE DATABASE {{ .DB.Name }} ENCODING 'UTF8';
CREATE ROLE {{ .DB.AppUser.Username }} LOGIN CREATEROLE;
CREATE USER {{ .DB.Admin.Username }} PASSWORD '{{ .DB.Admin.Password }}';
GRANT {{ .DB.AppUser.Username }} to {{ .DB.Admin.Username }};
GRANT ALL PRIVILEGES ON DATABASE {{ .DB.Name }} to {{ .DB.AppUser.Username }};
-- ALTER DATABASE {{ .DB.Name }} OWNER TO {{ .DB.AppUser.Username }};

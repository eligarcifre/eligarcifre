var (
	id      int
	nombre  string
	pass    string
	permiso int
)
//rows, err := db.Query("SELECT id, nombre, pass, permiso FROM public.usuarios WHERE id = ?", 1)
rows, err := db.Query("SELECT * FROM public.usuarios ORDER BY id ")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&id, &nombre, &pass, &permiso)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id, nombre, pass, permiso)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}


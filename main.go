func main() {
	router := gin.Default()

		db := pg.Connect(&pg.Options{
				User:     "postgres",
						Password: "mypassword",
								Database: "mydatabase",
									})
										defer db.Close()

											// quiz 테이블을 생성한다.
												createTableSQL := `
														CREATE TABLE IF NOT EXISTS quiz (
																	id serial PRIMARY KEY,
																				title TEXT NOT


package migrations

func IndexMigration() {
	CommentMigration()
	ArticleMigration()
	UserMigration()
}

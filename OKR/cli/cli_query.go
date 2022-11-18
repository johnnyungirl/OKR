package cli

import (
	"crud_app/database"
)

func (cli *CommandLine) list_user_okr() {
	db := database.Database{DbName: "okr"}
	db.Query_list_user()
}
func (cli *CommandLine) list_user_objs_okr() {
	db := database.Database{DbName: "okr"}
	db.Query_list_numsobj_user()
}

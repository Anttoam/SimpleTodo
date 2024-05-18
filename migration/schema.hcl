table "users" {
  schema = schema.main
  column "id" {
    null = true
    type = integer
  }
  column "name" {
    null = false
    type = text
  }
  column "email" {
    null = false
    type = text
  }
  column "password" {
    null = false
    type = text
  }
  column "created_at" {
    null = true
    type = sql("timestamp")
  }
  column "updated_at" {
    null = true
    type = sql("timestamp")
  }
  primary_key {
    columns = [column.id]
  }
  index "users_email" {
    unique  = true
    columns = [column.email]
  }
}
table "todos" {
  schema = schema.main
  column "id" {
    null = true
    type = integer
  }
  column "title" {
    null = false
    type = text
  }
  column "description" {
    null = false
    type = text
  }
  column "done" {
    null    = false
    type    = boolean
    default = false
  }
  column "user_id" {
    null = true
    type = text
  }
  column "created_at" {
    null = true
    type = sql("timestamp")
  }
  column "updated_at" {
    null = true
    type = sql("timestamp")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "0" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
schema "main" {
}

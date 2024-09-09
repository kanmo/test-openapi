table "users" {
  schema = schema.test
  column "id" {
    null = false
    type = int
  }
  column "name" {
    null = false
    type = varchar(100)
  }
  column "pw_hash" {
    null = false
    type = varchar(100)
  }
  primary_key {
    columns = [column.id]
  }
}

table "spaces" {
    schema = schema.test
    column "space_id" {
        null = false
        type = int
        auto_increment = true
    }
    column "name" {
        null = false
        type = varchar(100)
    }
    column "owner" {
        null = false
        type = varchar(100)
    }
    primary_key {
        columns = [column.space_id]
    }
}

table "messages" {
    schema = schema.test
    column "message_id" {
        null = false
        type = int
    }
    column "space_id" {
        null = false
        type = int
    }
    column "author" {
        null = false
        type = varchar(100)
    }
    column "content" {
        null = false
        type = varchar(1024)
    }
    column "timestamp" {
        null = false
        type = timestamp
        default = sql("CURRENT_TIMESTAMP")
    }
    primary_key {
        columns = [column.message_id]
    }
    foreign_key "fk_messages_space_id" {
        columns     = [column.space_id]
        ref_columns = [table.spaces.column.space_id]
        on_update   = CASCADE
        on_delete   = CASCADE
  }
}
schema "test" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}


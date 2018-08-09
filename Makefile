SRC_DIR 	= 	src/

SRC 		= 	main.go \
			basic_handlers.go \
			logger.go \
			pars_routes.go \
			router.go \
			routes.go \
			sql_db_utils.go \
			pars_cmdline.go \
			get_config.go \
			get_table_sql.go \
			get_line_sql.go \
			create_line_sql.go \
			handlers_utils.go \
			update_line_sql.go \
			delete_line_sql.go \
			redis_utils.go \
			add_entry_redis.go

SRC 		:= 	$(addprefix $(SRC_DIR), $(SRC))

NAME 		= 	rest_api

$(NAME): 		
			go build -o $(NAME) $(SRC)

all: 			$(NAME)

fclean: 		
			rm -f $(NAME)

re: 			fclean all

deps: 			
			go get github.com/gorilla/mux
			go get github.com/go-sql-driver/mysql
			go get github.com/mediocregopher/radix.v2

.PHONY: 		all re deps fclean

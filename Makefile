SRC_DIR 	= 	src/

SRC 		= 	main.go \
			handlers.go \
			logger.go \
			pars_routes.go \
			router.go \
			routes.go \
			todo.go \
			db_utils.go \
			pars_cmdline.go \
			get_db_login.go \
			get_table.go

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

.PHONY: 		all re deps fclean

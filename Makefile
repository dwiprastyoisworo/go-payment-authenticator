# Variabel
APP_NAME=myapp
CMD_PATH=./cmd
MAIN_FILE=app.serve.go

# Target untuk menjalankan aplikasi
run:
	go run $(CMD_PATH)/$(MAIN_FILE)

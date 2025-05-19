BIN_DIR := bin

build-windows:
	mkdir -p $(BIN_DIR)/win64
	env GOOS=windows GOARCH=amd64 \
		CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
		CGO_LDFLAGS="-lmingw32 -lSDL2" \
		go build -o $(BIN_DIR)/win64/craftgame.exe
	cp terrain.png $(BIN_DIR)/win64/terrain.png
	cp -r libs/win64/* $(BIN_DIR)/win64

build-linux:
	mkdir -p $(BIN_DIR)/linux64
	env GOOS=linux GOARCH=amd64 \
		go build -o $(BIN_DIR)/linux64/craftgame
	cp terrain.png $(BIN_DIR)/linux64/terrain.png

@echo off
echo [1/3] Building frontend...
cd frontend
call npm install
call npm run build
cd ..

echo [2/3] Downloading Go dependencies...
go mod tidy

echo [3/3] Building Go binary...
set CGO_ENABLED=1
go build -o giftlist.exe .

echo.
echo Build complete! Run with:
echo   set ADMIN_PASSWORD=your_password
echo   giftlist.exe

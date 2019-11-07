go clean
go build -o terraform-provider-clouddk.exe
copy terraform-provider-clouddk.exe %APPDATA%\terraform.d\plugins\windows_amd64\terraform-provider-clouddk.exe
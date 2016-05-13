set GOPATH=f:\work\apiserver
set PATH=%PATH%;f:\work\apiserver\bin

go get github.com/tools/godep

godep update

#godep get github.com/BurntSushi/toml github.com/gorilla/mux github.com/SermoDigital/jose github.com/tamnd/gauth github.com/tamnd/httpclient golang.org/x/oauth2 github.com/gokrokvertskhov/gauth

#godep save github.com/BurntSushi/toml github.com/gorilla/mux github.com/SermoDigital/jose github.com/tamnd/gauth github.com/tamnd/httpclient golang.org/x/oauth2 github.com/gokrokvertskhov/gauth

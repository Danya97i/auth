package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i UserRepository,LogRepository -o ./mocks/ -s "_minimock.go"

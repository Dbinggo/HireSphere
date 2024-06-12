package databases

type DataBase interface {
	GetDsn() string
	InitDataBases() error
}

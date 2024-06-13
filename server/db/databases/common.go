package databases

type DataBase interface {
	getDsn() string
	InitDataBases() error
}

package database

type _Getter interface {
	getter(field string)
}

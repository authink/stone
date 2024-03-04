package sql

type Inserter interface {
	Insert() string
}

type Saver interface {
	Save() string
}

type Deleter interface {
	Delete() string
}

type Updater interface {
	Update() string
}

type Geter interface {
	Get() string
}

type GeterForUpdate interface {
	GetForUpdate() string
}

type Finder interface {
	Find() string
}

type Counter interface {
	Count() string
}

type Pager interface {
	Pagination() string
}
